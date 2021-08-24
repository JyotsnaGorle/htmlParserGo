package htmlLinks

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

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
