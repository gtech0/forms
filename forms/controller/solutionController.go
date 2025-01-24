package controller

import "github.com/gin-gonic/gin"

type SolutionController struct{}

func NewSolutionController() *SolutionController {
	return &SolutionController{}
}

func (s *SolutionController) CreateSolution(ctx *gin.Context) {

}
