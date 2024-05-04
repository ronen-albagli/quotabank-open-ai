package web

import (
	"fmt"
	"net/http"

	core "genaidemo/pkg/core"

	"github.com/gin-gonic/gin"
)

type TranslationInput struct {
	WordToTranslate string `json:"wordToTranslate" binding:"required"`
}

type TranslationError struct {
	Error string `json:"error"`
}

type TranslationResponse struct {
	Translation string `json:"translation"`
}

func CreateTranslationHandler(c *gin.Context) {
	var parsedInput TranslationInput

	if err := c.BindJSON(&parsedInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation error", "details": err.Error()})
		return
	}

	translateResponse, err := core.Translate(parsedInput.WordToTranslate)

	if err != nil {
		c.JSON(http.StatusBadRequest, TranslationError{Error: err.Error()})
		return
	}

	fmt.Println(translateResponse)
	c.JSON(http.StatusCreated, TranslationResponse{Translation: translateResponse})
}
