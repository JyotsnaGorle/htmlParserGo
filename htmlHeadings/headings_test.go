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

	expectedResult := []HeadingsResult{
		{
			Heading: "h1",
			Count:   1},
		{
			Heading: "h2",
			Count:   4},
		{
			Heading: "h3",
			Count:   2},
		{
			Heading: "h4",
			Count:   0},
		{
			Heading: "h5",
			Count:   0},
		{
			Heading: "h6",
			Count:   0},
	}

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
