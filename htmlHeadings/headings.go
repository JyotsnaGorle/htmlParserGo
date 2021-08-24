package htmlHeadings

import (
	"github.com/PuerkitoBio/goquery"
)

func FindHeadings(doc goquery.Document) map[string]int {

	headingLevels := map[string]int{"h1": 0, "h2": 0, "h3": 0, "h4": 0, "h5": 0, "h6": 0}
	for heading, _ := range headingLevels {

		doc.Find(heading).Each(func(i int, s *goquery.Selection) {
			headingLevels[heading] = headingLevels[heading] + 1
		})
	}

	return headingLevels
}
