package controller

import "github.com/gin-gonic/gin"

type FormGeneratedController struct{}

func NewFormGeneratedController() *FormGeneratedController {
	return &FormGeneratedController{}
}

func (f *FormGeneratedController) GetMyForm(ctx *gin.Context) {
}
