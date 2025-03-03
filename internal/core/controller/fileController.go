package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/service"
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

// UploadFile godoc
// @Tags         File
// @Summary      Upload file
// @Description  upload file
// @Produce      json
// @Param   	 payload body get.FileDto false "File data"
// @Success      200 {object} minio.UploadInfo
// @Failure      400 {object} errs.CustomError
// @Router       /file/upload [post]
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

// DownloadFile godoc
// @Tags         File
// @Summary      Download file
// @Description  download file
// @Produce      json
// @Param   	 fileId path string true "File id"
// @Success      200
// @Failure      400 {object} errs.CustomError
// @Router       /file/download/{fileId} [get]
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
