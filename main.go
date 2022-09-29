package main

import (
	"go_google_images/googleImageCrawler"
)

func main() {

	query := "네이버 웨일"
	path := "./images"

	googleImageCrawler.Crawler(query, path)

}
