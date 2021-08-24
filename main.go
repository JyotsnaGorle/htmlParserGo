package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	headings "htmlParserGo/htmlHeadings"
	links "htmlParserGo/htmlLinks"
	version "htmlParserGo/htmlVersion"

	"github.com/PuerkitoBio/goquery"
)

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func pingURL(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	fmt.Print(res.Header.Get("!DOCTYPE"))
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	html, err := doc.Html()
	if err != nil {
		return
	}

	version := version.CheckDoctype(html)
	fmt.Print(version)

	links.FindLinks(*doc)
	headings.FindHeadings(*doc)

}

func main() {
	// check if valid URL.

	urlToProccess := "https://www.htmldog.com/guides/html/beginner/headings/"
	isValidUrl(urlToProccess)
	pingURL(urlToProccess)
	// check HTML version.
	// parse for headers in each level.
	// Amount of internal and external links
	// Amount of inaccessible links
	// If a page contains a login form
}

// func customRouteHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	switch r.Method {
// 	case "GET":
// 		http.ServeFile(w, r, "form.html")
// 	case "POST":
// 		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
// 		if err := r.ParseForm(); err != nil {
// 			fmt.Fprintf(w, "ParseForm() err: %v", err)
// 			return
// 		}
// 		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)

// 		urlToProccess := r.FormValue("urlToProccess")

// 		if isValid := isValidUrl(urlToProccess); isValid {
// 			pingURL(urlToProccess)
// 		}

// 		fmt.Fprintf(w, "urlToProccess = %s\n", urlToProccess)
// 	default:
// 		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
// 	}
// }

// func main() {
// 	http.HandleFunc("/", customRouteHandler)

// 	fmt.Printf("Starting server for testing HTTP POST...\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
