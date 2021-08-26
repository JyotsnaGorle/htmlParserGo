package htmlLogin

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/* FindLogins: Check if a html document has a login form:
   Param: doc (goquery.Document) html-document
   Returns: boolean
*/
func FindLogins(doc goquery.Document) bool {

	var foundLogin bool

	doc.Find("form").Each(func(_ int, s *goquery.Selection) {

		s.Find("input").Each(func(i int, s *goquery.Selection) {

			typ, _ := s.Attr("type")
			typ = strings.ToLower(typ)

			if typ == "password" || typ == "submit" {
				name, _ := s.Attr("name")
				if name == "password" {
					fmt.Printf("Login found, %s, %s", typ, name)
				}
				foundLogin = true
			}

		})
	})

	return foundLogin

}
