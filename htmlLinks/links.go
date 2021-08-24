package htmlLinks

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Links struct {
	Internal    int
	External    int
	Inaccesable int
}

func FindLinks(doc goquery.Document) Links {

	var internalLinks []string
	var externalLinks []string

	var invalidLinks []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		link, exists := s.Attr("href")

		if !exists {
			return
		}
		FindInternalAndExternalLinks(link, internalLinks, externalLinks)
	})

	FindInaccesibleLinks(externalLinks, invalidLinks)

	return Links{
		Internal:    len(internalLinks),
		External:    len(externalLinks),
		Inaccesable: len(invalidLinks),
	}

}

func FindInternalAndExternalLinks(urlToProccess string, internalLinks []string, externalLinks []string) {

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

func FindInaccesibleLinks(links []string, invalidLinks []string) {
	for _, link := range links {
		_, err := url.Parse(link)
		if err != nil {
			return
		}

		pingUrl(link, invalidLinks)
	}
}

func pingUrl(urlToPing string, invalidLinks []string) {
	res, err := http.Get(urlToPing)
	if err != nil {
		invalidLinks = append(invalidLinks, urlToPing)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		invalidLinks = append(invalidLinks, urlToPing)
	}
}
