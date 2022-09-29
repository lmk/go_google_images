package googleImageCrawler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Crawler Get images from google
func Crawler(query string, path string) {

	makeDir(path)

	doc := requestDocument(query)
	no := 0

	var wait sync.WaitGroup

	doc.Find("h1").Each(func(i int, s *goquery.Selection) {

		if s.Text() == "Search results" {
			s = s.Next().Find("div")
			s.Find("img").Each(func(i int, s *goquery.Selection) {
				imgUrl := ""
				if src, exists := s.Attr("src"); exists {
					imgUrl = src
				} else if src, exists := s.Attr("data-src"); exists {
					imgUrl = src
				}

				if !strings.HasPrefix(imgUrl, "http") {
					return
				}

				no++
				filename := fmt.Sprintf("%s/%s_%04d", path, query, no) // without ext

				wait.Add(1)
				go func(url string, filename string) {
					saveUrlImage(imgUrl, filename)
					defer wait.Done()
				}(imgUrl, filename)
			})
		}
	})

	wait.Wait()
}

func makeDir(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func requestDocument(query string) *goquery.Document {

	url := fmt.Sprintf("https://www.google.co.in/search?q=%s&source=lnms&tbm=isch", strings.ReplaceAll(query, " ", "+"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

// saveUrlImage download url image, filesname is without extension
func saveUrlImage(url string, filename string) {

	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	imageType := getImageType(response)
	if imageType == "" {
		return
	}

	filename = fmt.Sprintf("%s.%s", filename, imageType)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(filename)
}

func getImageType(res *http.Response) string {
	imageTypes := []string{"jpeg", "svg", "png", "gif", "bmp"}

	contentType := res.Header.Get("Content-Type")

	for _, imageType := range imageTypes {
		if fmt.Sprintf("image/%s", imageType) == contentType {
			return imageType
		}
	}

	return ""
}
