package htmlLinks

import (
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

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

func TestFindHtmlLinks(t *testing.T) {
	testUrl := "https://www.github.com"

	expectedResult := Links{
		Internal:    72,
		External:    44,
		Inaccesable: 1,
	}

	doc := pingURL(testUrl)
	result := FindLinks(*doc)

	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("Failed: expected links not found %d, %d, %d", result.Internal, result.External, result.Inaccesable)
		// TODO: log mismatch
	}
}

func TestFindLinksCount(t *testing.T) {

	testLink := "https://www.github.com"

	var internalLinks []string
	var externalLinks []string

	findInternalAndExternalLinks(testLink, &internalLinks, &externalLinks)

	if len(externalLinks) != 1 {
		t.Errorf("Failed: did not detect external link, %s", testLink)
	}
}