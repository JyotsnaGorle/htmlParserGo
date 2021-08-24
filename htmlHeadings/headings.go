package htmlHeadings

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func FindHeadings(doc goquery.Document) {

	headingLevels := map[string]int{"h1": 0, "h2": 0, "h3": 0, "h4": 0, "h5": 0, "h6": 0}
	for heading, headingCount := range headingLevels {

		selector := heading

		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			headingCount = headingCount + 1
		})
	}

	fmt.Print(headingLevels)
}
