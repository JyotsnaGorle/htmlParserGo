package htmlHeadings

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func FindHeadings(doc goquery.Document) {

	n := 5
	values := make([]int, n)
	for i := range values {

		selector := "h" + strconv.Itoa(i+1)

		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			values[i] = values[i] + 1
		})
	}

	fmt.Print(values)
}
