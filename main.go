package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	helpers "htmlParserGo/helpers"
	headings "htmlParserGo/htmlHeadings"
	links "htmlParserGo/htmlLinks"
	login "htmlParserGo/htmlLogin"
	version "htmlParserGo/htmlVersion"

	"github.com/PuerkitoBio/goquery"
)

type HtmlParseResult struct {
	UrlToParse string
	Version    string
	Title      string
	Headings   []headings.HeadingsResult
	Links      links.Links
	HasLogin   bool
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
		http.ServeFile(w, r, "./frontend/templates/form.html")

	case "POST":
		tmpl := template.Must(template.ParseFiles("./frontend/templates/layout.html"))

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		urlToProccess := r.FormValue("urlToProccess")

		finalResult := HtmlParseResult{
			UrlToParse: urlToProccess,
		}

		if isValid := helpers.IsValidUrl(urlToProccess); isValid {

			pingURL(urlToProccess, &finalResult)
			tmpl.Execute(w, finalResult)

			/*
				TO USE: in case of REST server application:
				-------------------------------------------
					w.Header().Set("Content-Type", "application/json")
						json.NewEncoder(w).Encode(finalResult)
					} else {
						http.Error(w, "400 invalid parameter value.", http.StatusBadRequest)
			*/

		} else {
			http.Error(w, "400 invalid parameter value.", http.StatusBadRequest)
		}

	default:
		fmt.Fprintf(w, "Only GET and POST methods are supported.")
	}
}

func main() {

	fs := http.FileServer(http.Dir("./frontend/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	js := http.FileServer(http.Dir("./frontend/js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	http.HandleFunc("/", customRouteHandler)

	fmt.Printf("Starting server at localhost:8000...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
