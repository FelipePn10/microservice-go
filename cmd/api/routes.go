package main

import (
	"github.com/FelipePn10/microservice-go/cmd/api/controllers"
	"github.com/FelipePn10/microservice-go/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	router := gin.Default()

	categoryRoutes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()
	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

}
