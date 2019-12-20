package models

// QuoteModel structure holds the author and quote
type QuoteModel struct {
	Quote       string `json:"quote"`
	Author      string `json:"author"`
	Photo       string `json:"photo"`
	PhotoAuthor string `json:"photoAuthor"`
}
