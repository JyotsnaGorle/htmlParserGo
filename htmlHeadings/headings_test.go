package htmlHeadings

import (
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestHtmlHeadings(t *testing.T) {

	testUrl := "https://www.htmldog.com/guides/html/beginner/headings/"

	expectedResult := map[string]int{"h1": 1, "h2": 4, "h3": 2, "h4": 0, "h5": 0, "h6": 0}

	doc := pingURL(testUrl)
	result := FindHeadings(*doc)

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("Failed: expected headers not found")
		// TODO: log mismatch
	}
}

func pingURL(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
