package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/mapper"
	"hedgehog-forms/internal/core/service"
	"net/http"
)

type FormPublishedController struct {
	formPublishedMapper  *mapper.FormPublishedMapper
	formPublishedService *service.FormPublishedService
}

func NewFormPublishedController() *FormPublishedController {
	return &FormPublishedController{
		formPublishedMapper:  mapper.NewFormPublishedMapper(),
		formPublishedService: service.NewFormPublishedService(),
	}
}

// PublishForm godoc
// @Tags         FormPublished
// @Summary      Publish form
// @Description  publish form
// @Produce      json
// @Param   	 payload body create.FormPublishDto false "Form data"
// @Success      200 {object} get.FormPublishedBaseDto
// @Failure      400 {object} errs.CustomError
// @Router       /form/published/create [post]
func (f *FormPublishedController) PublishForm(ctx *gin.Context) {
	var publishDto create.FormPublishDto
	if err := ctx.Bind(&publishDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	formPublished, err := f.formPublishedService.PublishForm(publishDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formPublished)
}

// GetForm godoc
// @Tags         FormPublished
// @Summary      Get form
// @Description  get form
// @Produce      json
// @Param   	 publishedId path string true "Published id"
// @Success      200 {object} get.FormPublishedDto
// @Failure      400 {object} errs.CustomError
// @Router       /form/published/get/{publishedId} [get]
func (f *FormPublishedController) GetForm(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")

	formPublished, err := f.formPublishedService.GetForm(publishedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formPublished)
}

func (f *FormPublishedController) GetForms(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	forms, err := f.formPublishedService.GetForms(query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *forms)
}

// UpdateForm godoc
// @Tags         FormPublished
// @Summary      Update form
// @Description  update form
// @Produce      json
// @Param   	 payload body create.UpdateFormPublishedDto false "New form data"
// @Param   	 publishedId path string true "Published id"
// @Success      200 {object} get.FormPublishedDto
// @Failure      400 {object} errs.CustomError
// @Router       /form/published/update/{publishedId} [put]
func (f *FormPublishedController) UpdateForm(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	var updateFormPublishedDto create.UpdateFormPublishedDto
	if err := ctx.Bind(&updateFormPublishedDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	response, err := f.formPublishedService.UpdateForm(publishedId, updateFormPublishedDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *response)
}

// DeleteForm godoc
// @Tags         FormPublished
// @Summary      Delete form
// @Description  delete form
// @Produce      json
// @Param   	 publishedId path string true "Published id"
// @Success      200
// @Failure      400 {object} errs.CustomError
// @Router       /form/published/delete/{publishedId} [delete]
func (f *FormPublishedController) DeleteForm(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	err := f.formPublishedService.DeleteForm(publishedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
