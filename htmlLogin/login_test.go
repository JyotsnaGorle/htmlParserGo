package htmlLogin

import (
	"log"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestLoginDoesNotExists(t *testing.T) {

	testUrl := "https://www.stealmylogin.com"
	doc := pingURL(testUrl)
	if FindLogins(*doc) {
		t.Errorf("Failed: login found %s", testUrl)
	}

}

func TestLoginExists(t *testing.T) {

	testUrl := "https://www.stealmylogin.com/demo.html"
	doc := pingURL(testUrl)
	if !FindLogins(*doc) {
		t.Errorf("Failed: login not found %s", testUrl)
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
