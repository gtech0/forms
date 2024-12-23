package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/service"
	"io"
	"net/http"
	"os"
)

type FileController struct {
	bucket      string
	fileService *service.FileService
}

func NewFileController() *FileController {
	return &FileController{
		bucket:      os.Getenv("MINIO_BUCKET"),
		fileService: service.NewFileService(),
	}
}

func (f *FileController) UploadFile(ctx *gin.Context) {
	fileDto := get.FileDto{}
	if err := ctx.Bind(&fileDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	info, err := f.fileService.UploadFile(fileDto, f.bucket)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *info)
}

func (f *FileController) DownloadFile(ctx *gin.Context) {
	fileId := ctx.Param("fileId")
	reader, err := f.fileService.DownloadFile(fileId, f.bucket)
	if err != nil {
		ctx.Error(err)
		return
	}

	info, err := reader.Stat()
	if err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	ctx.Header("Content-Disposition", info.Metadata.Get("Content-Disposition"))
	ctx.Header("Content-Type", info.Metadata.Get("Content-Type"))
	if _, err = io.Copy(ctx.Writer, reader); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	ctx.Status(http.StatusOK)
}
