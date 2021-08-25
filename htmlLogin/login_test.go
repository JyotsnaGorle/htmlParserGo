package htmlLogin

import (
	"log"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestLoginDoesNotExists(t *testing.T) {

	testUrl := "https://www.stealmylogin.com"

	doc, err := pingURL(testUrl)

	if err != nil {
		t.Errorf("Failed: could not ping url")
	}

	if FindLogins(*doc) {
		t.Errorf("Failed: login found %s", testUrl)
	}

}

func TestLoginExists(t *testing.T) {

	testUrl := "https://www.stealmylogin.com/demo.html"

	doc, err := pingURL(testUrl)

	if err != nil {
		t.Errorf("Failed: could not ping url")
	}

	if !FindLogins(*doc) {
		t.Errorf("Failed: login not found %s", testUrl)
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
