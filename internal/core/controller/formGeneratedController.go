package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/mapper"
	"hedgehog-forms/internal/core/service"
	"net/http"
)

type FormGeneratedController struct {
	formGeneratedService *service.FormGeneratedService
	formGeneratedMapper  *mapper.FormGeneratedMapper
}

func NewFormGeneratedController() *FormGeneratedController {
	return &FormGeneratedController{
		formGeneratedService: service.NewFormGeneratedService(),
		formGeneratedMapper:  mapper.NewFormGeneratedMapper(),
	}
}

// Create godoc
// @Tags         FormGenerated
// @Summary      Create user form
// @Description  Create user form
// @Produce      json
// @Param   	 publishedId path string true "Published id"
// @Success      200 {object} get.MyGeneratedDto
// @Failure      400 {object} errs.CustomError
// @Router       /form/generated/create/{publishedId} [post]
func (f *FormGeneratedController) Create(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	var userIdDto create.FormGeneratedUser
	if err := ctx.Bind(&userIdDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	formGeneratedDto, err := f.formGeneratedService.Create(publishedId, userIdDto.UserId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formGeneratedDto)
}

func (f *FormGeneratedController) Get(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")
	formGeneratedDto, err := f.formGeneratedService.Get(generatedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formGeneratedDto)
}

// GetMyForm godoc
// @Tags         FormGenerated
// @Summary      Get current user form
// @Description  get current user form
// @Produce      json
// @Param   	 publishedId path string true "Published id"
// @Success      200 {object} get.FormGeneratedDto
// @Failure      400 {object} errs.CustomError
// @Router       /form/generated/get/{publishedId} [post]
//func (f *FormGeneratedController) GetMyForm(ctx *gin.Context) {
//	publishedId := ctx.Param("publishedId")
//	var userIdDto create.FormGeneratedUser
//	if err := ctx.Bind(&userIdDto); err != nil {
//		ctx.Error(errs.New(err.Error(), 500))
//		return
//	}
//
//	formGeneratedDto, err := f.formGeneratedService.GetMyForm(publishedId, userIdDto.UserId)
//	if err != nil {
//		ctx.Error(err)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, *formGeneratedDto)
//}

// SaveAnswers godoc
// @Tags         FormGenerated
// @Summary      Save answers for current form
// @Description  save answers for current form
// @Produce      json
// @Param   	 generatedId path string true "Generated id"
// @Success      200 {object} get.FormGeneratedDto
// @Failure      400 {object} errs.CustomError
// @Router       /form/generated/save/{generatedId} [post]
//func (f *FormGeneratedController) SaveAnswers(ctx *gin.Context) {
//	generatedId := ctx.Param("generatedId")
//	var answerDto get.AnswerDto
//	if err := ctx.Bind(&answerDto); err != nil {
//		ctx.Error(errs.New(err.Error(), 500))
//		return
//	}
//
//	formGeneratedDto, err := f.formGeneratedService.SaveAnswers(generatedId, answerDto)
//	if err != nil {
//		ctx.Error(err)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, *formGeneratedDto)
//}

// SubmitForm godoc
// @Tags         FormGenerated
// @Summary      Submit current form for evaluation
// @Description  submit current form for evaluation
// @Produce      json
// @Param   	 generatedId path string true "Generated id"
// @Success      200 {object} get.MyGeneratedDto
// @Failure      400 {object} errs.CustomError
// @Router       /form/generated/submit/{generatedId} [post]
//func (f *FormGeneratedController) SubmitForm(ctx *gin.Context) {
//	generatedId := ctx.Param("generatedId")
//	var answerDto get.AnswerDto
//	if err := ctx.Bind(&answerDto); err != nil {
//		ctx.Error(errs.New(err.Error(), 500))
//		return
//	}
//
//	myGeneratedDto, err := f.formGeneratedService.SubmitForm(generatedId, answerDto)
//	if err != nil {
//		ctx.Error(err)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, *myGeneratedDto)
//}

func (f *FormGeneratedController) UnSubmitForm(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")

	myGeneratedDto, err := f.formGeneratedService.UnSubmitForm(generatedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *myGeneratedDto)
}

func (f *FormGeneratedController) GetMyForms(ctx *gin.Context) {
	subjectId := ctx.Param("subjectId")
	userId := ctx.Param("userId")
	query := ctx.Request.URL.Query()
	response, err := f.formGeneratedService.GetMyForms(userId, subjectId, query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *response)
}

func (f *FormGeneratedController) GetSubmittedForms(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	userId := ctx.Param("userId")
	query := ctx.Request.URL.Query()
	response, err := f.formGeneratedService.GetSubmittedForms(userId, publishedId, query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *response)
}

//func (f *FormGeneratedController) GetSubmittedForm(ctx *gin.Context) {
//	generatedId := ctx.Param("generatedId")
//	submittedForm, err := f.formGeneratedService.GetSubmittedForm(generatedId)
//	if err != nil {
//		ctx.Error(err)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, *submittedForm)
//}

// VerifyForm godoc
// @Tags         FormGenerated
// @Summary      Evaluate users form manually
// @Description  evaluate users form manually
// @Produce      json
// @Param   	 generatedId path string true "Generated id"
// @Success      200 {object} get.FormGeneratedDto
// @Failure      400 {object} errs.CustomError
// @Router       /form/generated/verify/{generatedId} [post]
func (f *FormGeneratedController) VerifyForm(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")
	var checkDto create.CheckDto
	if err := ctx.Bind(&checkDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	formGeneratedDto, err := f.formGeneratedService.VerifyForm(generatedId, checkDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formGeneratedDto)
}

func (f *FormGeneratedController) ReturnForm(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")

	myGeneratedDto, err := f.formGeneratedService.ReturnForm(generatedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *myGeneratedDto)
}
