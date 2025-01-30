package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/file"
	"hedgehog-forms/internal/core/model"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
	"hedgehog-forms/internal/core/repository"
	"mime/multipart"
)

type FileService struct {
	fileRepository *repository.FileRepository
}

func NewFileService() *FileService {
	return &FileService{
		fileRepository: repository.NewFileRepository(),
	}
}

func (f *FileService) UploadFile(fileDto get.FileDto, bucket string) (*minio.UploadInfo, error) {
	fileEntity := fileDto.File
	headers := fileEntity.Header

	fileModel := new(model.File)
	fileModel.Id = uuid.New()
	fileModel.Bucket = bucket
	fileModel.Size = fileEntity.Size

	fileOpened, err := fileEntity.Open()
	defer func(fileOpened multipart.File) {
		err = fileOpened.Close()
	}(fileOpened)

	if err != nil {
		return nil, errs.New(err.Error(), 500)
	}

	options := minio.PutObjectOptions{
		ContentType:        headers.Get("Content-Type"),
		ContentDisposition: headers.Get("Content-Disposition"),
	}

	info, err := file.MinioClient.PutObject(
		context.Background(), bucket, fileModel.Id.String(), fileOpened, fileModel.Size, options,
	)
	if err != nil {
		return nil, errs.New(err.Error(), 500)
	}

	if err = f.fileRepository.Save(fileModel); err != nil {
		return nil, err
	}

	return &info, nil
}

func (f *FileService) DownloadFile(fileId string, bucket string) (*minio.Object, error) {
	if fileId == "" {
		return nil, errs.New("fileId is required", 400)
	}

	reader, err := file.MinioClient.GetObject(context.Background(), bucket, fileId, minio.GetObjectOptions{})
	if err != nil {
		return nil, errs.New(err.Error(), 500)
	}

	defer func(reader *minio.Object) {
		err = reader.Close()
	}(reader)

	if err != nil {
		return nil, errs.New(err.Error(), 500)
	}

	return reader, nil
}

func (f *FileService) ValidateFiles(questions ...*question.Question) error {
	fileIds := make([]uuid.UUID, 0)
	for _, questionEntity := range questions {
		for _, attachment := range questionEntity.Attachments {
			fileIds = append(fileIds, attachment.FileId)
		}
	}

	nonexistentFileIds := make([]uuid.UUID, 0)
	for _, fileId := range fileIds {
		fileEntity, err := f.fileRepository.FindById(fileId)
		if err != nil {
			return err
		}

		if fileEntity.Id == uuid.Nil {
			nonexistentFileIds = append(nonexistentFileIds, fileId)
		}
	}

	if len(nonexistentFileIds) > 0 {
		return errs.New(fmt.Sprintf("incorrect attachment ids: %v", nonexistentFileIds), 400)
	}

	return nil
}
