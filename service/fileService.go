package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/file"
	"hedgehog-forms/model"
	"hedgehog-forms/repository"
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
	fileObj := fileDto.File
	headers := fileObj.Header

	fileModel := new(model.File)
	fileModel.Id = uuid.New()
	fileModel.Bucket = bucket
	fileModel.Size = fileObj.Size

	fileOpened, err := fileObj.Open()
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

	if err = f.fileRepository.Create(fileModel); err != nil {
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
