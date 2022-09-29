package googleImageCrawler

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestEnglish(t *testing.T) {
	testAndDel("IU", "testIU", t)
}

func TestKorean(t *testing.T) {
	testAndDel("아이유", "testIU2", t)
}

func TestContainSpace(t *testing.T) {
	testAndDel("안구 정화", "nice", t)
}

func testAndDel(query string, path string, t *testing.T) {
	Crawler(query, path)

	if !vaildFiles(path) {
		t.Error("contains size zero file")
	}

	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println(err)
	}
}

func vaildFiles(path string) bool {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}

	if len(files) == 0 {
		return false
	}

	sizeZero := false
	for _, file := range files {

		if file.Size() == 0 {
			sizeZero = true
			break
		}
	}

	return !sizeZero
}
