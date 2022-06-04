package server

import (
	"net/http"
	"reciper/recipe-generator/internal/env"
	"reciper/recipe-generator/internal/generator"
	"reciper/recipe-generator/internal/infrastructure"

	"github.com/gin-gonic/gin"
)

type input struct {
	Username string `json:"username"`
}

type output struct {
	Message string `json:"message"`
}

func CreateServer() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.POST("/generate", generateRecipe)

	serverPort := env.GetEnv().ServerPort

	router.Run("localhost:" + serverPort)
}

func ping(c *gin.Context) {
	output_data := output{
		Message: "pong",
	}
	c.JSON(http.StatusOK, output_data)
}

func generateRecipe(c *gin.Context) {
	var input_data input
	if err := c.BindJSON(&input_data); err != nil {
		output_data := output{
			Message: "invalid input data",
		}
		c.JSON(http.StatusBadRequest, output_data)
		return
	}
	recipe := generator.GenerateRecipe(input_data.Username)
	infrastructure.ProduceRecipeCreateEvent(recipe)
	c.JSON(http.StatusCreated, recipe)
}
