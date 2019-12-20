package storage

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"quote-me/config"
	"time"

	"github.com/valyala/fastjson"
)

// Choose a random keyword from an array of keywords and then search for it, choose the first image

func getRandomKeyword() string {

	rand.Seed(time.Now().Unix())

	keywords := config.GetConfig().GetStringSlice("unsplash.keywords")

	randomKeyword := keywords[rand.Intn(len(keywords))]

	return randomKeyword
}

// GetPhoto method returns a random photo from unsplash
func GetPhoto() (string, string) {

	randomKeyword := getRandomKeyword()

	resp, err := http.Get(fmt.Sprintf("https://api.unsplash.com/search/photos?per_page=10&client_id=%s&query=%s",
		config.GetConfig().GetString("unsplash.clientId"),
		randomKeyword))

	if err != nil {
		fmt.Println("Error occured while retrieving a random image from unsplash", err)
		return "", ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Cannot read the body", err)
	}

	parser := new(fastjson.Parser)
	v, _ := parser.Parse(string(body))
	//Use a random index from 0 until the lenght of the results
	randomPhotoIndex := rand.Intn(10)
	randomPhoto := v.GetArray("results")[randomPhotoIndex]
	photo, author := randomPhoto.GetStringBytes("urls", "small"), randomPhoto.GetStringBytes("user", "name")

	return string(photo), string(author)
}
