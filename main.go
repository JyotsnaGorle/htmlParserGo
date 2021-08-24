package main

import (
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
	login.FindLogins(*doc)

}

func main() {
	// https://www.stealmylogin.com/demo.html

	// "https://www.htmldog.com/guides/html/beginner/headings/"
	urlToProccess := "https://github.com/"
	helpers.IsValidUrl(urlToProccess)
	pingURL(urlToProccess)
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
