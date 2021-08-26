package main

import (
	"fmt"
	"html/template"
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
	UrlToParse string
	Version    string
	Title      string
	Headings   []headings.HeadingsResult
	Links      links.Links
	HasLogin   bool
}

type parseError struct {
	ErrorMsg string
}

func pingURL(urlToProccess string, finalResult *HtmlParseResult) error {

	// Fetch the url response
	res, err := http.Get(urlToProccess)
	if err != nil {
		return err
	}

	// close body at the end
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the reponse body
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	// Load the HTML document
	html, err := doc.Html()
	if err != nil {
		return err
	}

	finalResult.Version = version.CheckDoctype(html)
	finalResult.Title = doc.Find("title").Text()
	finalResult.Headings = headings.FindHeadings(*doc)
	finalResult.Links = links.FindLinks(*doc)
	finalResult.HasLogin = login.FindLogins(*doc)

	return nil
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

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		urlToProccess := r.FormValue("urlToProccess")

		finalResult := HtmlParseResult{
			UrlToParse: urlToProccess,
		}

		if isValid := helpers.IsValidUrl(urlToProccess); isValid {

			if err := pingURL(urlToProccess, &finalResult); err != nil {
				/*
					TO USE: in case of REST server application:
					-------------------------------------------
					http.Error(w, "Error reaching link: "+err.Error()+".", http.StatusBadRequest)
				*/

				errResult := parseError{

					ErrorMsg: err.Error(),
				}

				tmpl := template.Must(template.ParseFiles("./frontend/templates/error.html"))
				tmpl.Execute(w, errResult)

			} else {
				tmpl := template.Must(template.ParseFiles("./frontend/templates/layout.html"))
				tmpl.Execute(w, finalResult)

			}

			/*
				TO USE: in case of REST server application:
				-------------------------------------------
					w.Header().Set("Content-Type", "application/json")
						json.NewEncoder(w).Encode(finalResult)
					} else {
						http.Error(w, "400 invalid parameter value.", http.StatusBadRequest)
			*/

		} else {
			errResult := parseError{

				ErrorMsg: "400. Sorry, that is an invalid parameter value.",
			}

			tmpl := template.Must(template.ParseFiles("./frontend/templates/error.html"))
			tmpl.Execute(w, errResult)

			/*
				TO USE: in case of REST server application:
				-------------------------------------------
				http.Error(w, "400 invalid parameter value.", http.StatusBadRequest)
			*/
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

	fmt.Printf("Starting server at localhost:8080...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("****************")

		log.Fatal(err)
	}
}
