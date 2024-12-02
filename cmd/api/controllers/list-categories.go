package controllers

import (
	"net/http"

	"github.com/FelipePn10/microservice-go/internal/repositories"
	usecases "github.com/FelipePn10/microservice-go/internal/use-cases"
	"github.com/gin-gonic/gin"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCase := usecases.NewListCategoriesUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(400,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success":    true,
		"categories": categories,
	})
}
