## HTML Parser using Golang and Go-templates

A simple web application used to parse a url for meta information in the html page.

## Software specifications
    * Runs a Golang server in the backend, and Golang templates for frontend.
	* Recommended IDE Visual Studio Code.
    * Installed Golang in the machine. 
    * Use the following [link](https://golang.org/doc/install) for Golang installation.

## Build and Run the application
	* In order to build the project: type the command ``go build main.go`` in the root folder via the terminal.
	* To run the built binary, use ``go run main.go`` in the root folder via the terminal.
    * The application runs server in ``localhost:8000`` which can then be accessed by any browser.

## Sample output

    Link: https://www.github.com/
    Html version: HTML 5
    Title : GitHub: Where the world builds software · GitHub
    ----------------------------------------------------------
    Heading count:
        - h2: 19
        - h3: 22
        - h4: 7
        - h5: 0
        - h6: 0
        - h1: 1
    ----------------------------------------------------------
    Links:
        - Internal links count: 72
        - External links count: 44
        - Inaccesable links count: 1
    
    ----------------------------------------------------------
    No Login form found
    ----------------------------------------------------------