package htmlHeadings

import (
	"log"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestHtmlHeadings(t *testing.T) {

	testUrl := "https://www.htmldog.com/guides/html/beginner/headings/"

	expectedResult := map[string]int{"h1": 1, "h2": 4, "h3": 2, "h4": 0, "h5": 0, "h6": 0}

	doc, err := pingURL(testUrl)

	if err != nil {
		t.Errorf("Failed: could not ping url")
	}
	result := FindHeadings(*doc)

	for _, each := range result {
		if expectedResult[each.Heading] != each.Count {
			t.Errorf("Failed: expected headers not found for %s - expected : %d, actual : %d", each.Heading, expectedResult[each.Heading], each.Count)
		}
	}
}

func pingURL(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
