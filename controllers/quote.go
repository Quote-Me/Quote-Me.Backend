package controllers

import (
	"net/http"
	"quote-me/storage"

	"github.com/gin-gonic/gin"
)

// QuoteController for retrieving quotes
type QuoteController struct{}

// Quote method returns a random quote
func (q *QuoteController) Quote(c *gin.Context) {
	quote := storage.GetStorage().Quote()
	if quote == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No results found"})
	} else {
		c.JSON(http.StatusOK, quote)
	}
}
