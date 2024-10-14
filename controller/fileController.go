package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/file"
	"hedgehog-forms/model"
	"mime/multipart"
	"net/http"
	"os"
)

type FileController struct {
	bucket string
}

func NewFileController() *FileController {
	return &FileController{
		bucket: os.Getenv("MINIO_BUCKET"),
	}
}

func (f *FileController) UploadFile(ctx *gin.Context) {
	fileDto := get.FileDto{}
	if err := ctx.Bind(&fileDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileObj := fileDto.File
	headers := fileObj.Header

	fileModel := new(model.File)
	fileModel.Id = uuid.New()
	fileModel.Bucket = f.bucket
	fileModel.Size = fileObj.Size

	fileOpened, err := fileObj.Open()
	defer func(fileOpened multipart.File) {
		err = fileOpened.Close()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}(fileOpened)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	options := minio.PutObjectOptions{
		ContentType:        headers.Get("Content-Type"),
		ContentDisposition: headers.Get("Content-Disposition"),
	}

	info, err := file.MinioClient.PutObject(context.Background(), f.bucket,
		fileModel.Id.String(), fileOpened, fileModel.Size, options)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = database.DB.Create(&fileModel).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, info)
}

func (f *FileController) DownloadFile(ctx *gin.Context) {
	fileId := ctx.Param("fileId")
	if fileId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "fileId is required",
		})
		return
	}

	reader, err := file.MinioClient.GetObject(context.Background(), f.bucket, fileId, minio.GetObjectOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer func(reader *minio.Object) {
		err = reader.Close()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}(reader)

	stat, err := reader.Stat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, stat)
}
