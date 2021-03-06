package htmlLinks

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Links struct {
	Internal    int
	External    int
	Inaccesable int
}

/* FindLinks: In a html document, finds
   Internal, External and Inaccesable links.
   Param: doc (goquery.Document) html-document
   Returns: Links
*/
func FindLinks(doc goquery.Document) Links {

	var internalLinks []string
	var externalLinks []string

	var invalidLinks []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		link, exists := s.Attr("href")

		if !exists {
			return
		}
		findInternalAndExternalLinks(link, &internalLinks, &externalLinks)
	})

	findInaccesibleLinks(externalLinks, &invalidLinks)

	return Links{
		Internal:    len(internalLinks),
		External:    len(externalLinks),
		Inaccesable: len(invalidLinks),
	}

}

func findInternalAndExternalLinks(urlToProccess string, internalLinks *[]string, externalLinks *[]string) {

	u, err := url.Parse(urlToProccess)
	if err != nil {
		return
	}

	if u.Scheme == "" && u.Host == "" {
		*internalLinks = append(*internalLinks, urlToProccess)
	} else {
		*externalLinks = append(*externalLinks, urlToProccess)
	}
}

func findInaccesibleLinks(links []string, invalidLinks *[]string) {

	concurrencyLimit := 2

	if len(links) > 2 {
		concurrencyLimit = len(links) / 2
	}

	gaurd := make(chan struct{}, concurrencyLimit)
	wg := sync.WaitGroup{}

	// check links concurrently
	for _, link := range links {
		_, err := url.Parse(link)
		if err != nil {
			return
		}

		gaurd <- struct{}{}

		wg.Add(1)

		go func(urlLink string) {

			defer wg.Done()
			pingUrl(urlLink, invalidLinks)

			<-gaurd

		}(link)
	}
	wg.Wait()
}

// ping urls to check accesibility
func pingUrl(urlToPing string, invalidLinks *[]string) {

	res, err := http.Get(urlToPing)

	if err != nil {
		*invalidLinks = append(*invalidLinks, urlToPing)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		*invalidLinks = append(*invalidLinks, urlToPing)
	}
}
