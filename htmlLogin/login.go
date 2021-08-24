package htmlLogin

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Amount of internal and external links
// Amount of inaccessible links
func FindLogins(doc goquery.Document) {
	// Find the review items
	doc.Find("form").Each(func(_ int, s *goquery.Selection) {

		s.Find("input").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			name, _ := s.Attr("name")
			if name == "password" {
				fmt.Printf("Login found, %s", name)
			}

			typ, _ := s.Attr("type")
			typ = strings.ToLower(typ)
			fmt.Print(typ)

			if typ == "password" || typ == "submit" {
				fmt.Printf("Login found, %s", typ)
			}

		})
	})

}
