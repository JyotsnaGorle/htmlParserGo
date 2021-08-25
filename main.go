package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	helpers "htmlParserGo/helpers"
	headings "htmlParserGo/htmlHeadings"
	links "htmlParserGo/htmlLinks"
	login "htmlParserGo/htmlLogin"
	version "htmlParserGo/htmlVersion"

	"github.com/PuerkitoBio/goquery"
)

type HtmlParseResult struct {
	Version  string
	Title    string
	Headings []headings.HeadingsResult
	Links    links.Links
	HasLogin bool
}

func pingURL(urlToProccess string, finalResult *HtmlParseResult) {

	res, err := http.Get(urlToProccess)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	html, err := doc.Html()
	if err != nil {
		return
	}

	finalResult.Version = version.CheckDoctype(html)
	finalResult.Title = doc.Find("title").Text()
	finalResult.Headings = headings.FindHeadings(*doc)
	finalResult.Links = links.FindLinks(*doc)
	finalResult.HasLogin = login.FindLogins(*doc)

}

func customRouteHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		finalResult := HtmlParseResult{}
		urlToProccess := r.FormValue("urlToProccess")

		if isValid := helpers.IsValidUrl(urlToProccess); isValid {
			pingURL(urlToProccess, &finalResult)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(finalResult)
		} else {
			http.Error(w, "400 invalid parameter value.", http.StatusBadRequest)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", customRouteHandler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
