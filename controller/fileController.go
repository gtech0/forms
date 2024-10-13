package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/file"
	"mime/multipart"
	"net/http"
	"os"
)

type FileController struct{}

func NewFileController() *FileController {
	return &FileController{}
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

	bucket := os.Getenv("MINIO_BUCKET")
	info, err := file.MinioClient.PutObject(
		context.Background(),
		bucket,
		fileObj.Filename,
		fileOpened,
		fileObj.Size,
		minio.PutObjectOptions{
			ContentType:        headers.Get("Content-Type"),
			ContentDisposition: headers.Get("Content-Disposition"),
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, info)
}
