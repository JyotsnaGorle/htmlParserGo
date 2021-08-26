package htmlHeadings

import (
	"github.com/PuerkitoBio/goquery"
)

type HeadingsResult struct {
	Heading string
	Count   int
}

/* FindHeadings: Find frequency of all heading levels in html document.
   Param: doc (goquery.Document) html-document
   Returns: list of []HeadingResult.
*/
func FindHeadings(doc goquery.Document) []HeadingsResult {

	headingLevels := map[string]int{"h1": 0, "h2": 0, "h3": 0, "h4": 0, "h5": 0, "h6": 0}

	for heading, _ := range headingLevels {

		doc.Find(heading).Each(func(i int, s *goquery.Selection) {
			headingLevels[heading] = headingLevels[heading] + 1
		})
	}

	return fillStruct(headingLevels)
}

func fillStruct(mapData map[string]int) (hResult []HeadingsResult) {

	for k, v := range mapData {

		r := HeadingsResult{
			Heading: k,
			Count:   v,
		}

		hResult = append(hResult, r)
	}

	return hResult
}
