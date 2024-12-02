package controllers

import (
	"github.com/FelipePn10/microservice-go/internal/repositories"
	usecases "github.com/FelipePn10/microservice-go/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body CreateCategoryInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(400,
			gin.H{
				"success": true,
				"error":   err.Error(),
			})
		return
	}

	useCase := usecases.NewCreateCategoryUseCase(repository)
	err := useCase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(400,
			gin.H{
				"success": true,
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Category created successfully",
	})
}
