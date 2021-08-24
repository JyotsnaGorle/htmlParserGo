package htmlLogin

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FindLogins(doc goquery.Document) bool {

	var foundLogin bool

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
				foundLogin = true
			}

		})
	})

	return foundLogin

}
