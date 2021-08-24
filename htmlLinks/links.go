package htmlLinks

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

var internalLinks []string
var externalLinks []string

var invalidLinks []string

// Amount of internal and external links
// Amount of inaccessible links
func FindLinks(doc goquery.Document) {
	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		link, exists := s.Attr("href")

		if !exists {
			return
		}

		fmt.Printf("Review %d: %s\n", i, link)
	})
}

func FindInternalAndExternalLinks(urlToProccess string) {
	u, err := url.Parse(urlToProccess)
	if err != nil {
		return
	}

	if u.Scheme == "" && u.Host == "" {
		internalLinks = append(internalLinks, urlToProccess)
	} else {
		externalLinks = append(externalLinks, urlToProccess)
	}

}

func FindInaccesibleLinks(links []string) {
	for _, link := range links {
		_, err := url.Parse(link)
		if err != nil {
			return
		}

		pingUrl(link)
	}
}

func pingUrl(urlToPing string) {
	res, err := http.Get(urlToPing)
	if err != nil {
		invalidLinks = append(invalidLinks, urlToPing)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		invalidLinks = append(invalidLinks, urlToPing)
	}
}
