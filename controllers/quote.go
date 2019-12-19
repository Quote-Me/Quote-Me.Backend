package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// QuoteController for retrieving quotes
type QuoteController struct{}

// Quote method returns a random quote
func (q *QuoteController) Quote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
