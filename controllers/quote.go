package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"quote-me/storage"

	"github.com/gin-gonic/gin"
)

// QuoteController for retrieving quotes
type QuoteController struct{}

// Quote method returns a random quote
func (q *QuoteController) Quote(c *gin.Context) {
	quote := storage.GetStorage().Quote()
	quote.Photo, quote.PhotoAuthor = storage.GetPhoto()
	if quote == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No results found"})
	} else {
		//Use json.Marshal because gin escapes special characters
		buffer := bytes.NewBuffer([]byte{})
		enc := json.NewEncoder(buffer)
		enc.SetEscapeHTML(false)
		enc.Encode(quote)
		c.String(http.StatusOK, buffer.String())
	}
}
