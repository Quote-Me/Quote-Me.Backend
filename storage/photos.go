package storage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"quote-me/config"

	"github.com/valyala/fastjson"
)

// GetPhoto method returns a random photo from unsplash
func GetPhoto() string {
	resp, err := http.Get(fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%s", config.GetConfig().GetString("unsplash.clientId")))

	if err != nil {
		fmt.Println("Error occured while retrieving a random image from unsplash", err)
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Cannot read the body", err)
	}

	return fastjson.GetString([]byte(body), "urls", "small")
}
