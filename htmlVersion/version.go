package htmlVersion

import "strings"

var doctypes = make(map[string]string)

func init() {
	doctypes["HTML 4.01 Strict"] = `"-//W3C//DTD HTML 4.01//EN"`
	doctypes["HTML 4.01 Transitional"] = `"-//W3C//DTD HTML 4.01 Transitional//EN"`
	doctypes["HTML 4.01 Frameset"] = `"-//W3C//DTD HTML 4.01 Frameset//EN"`
	doctypes["XHTML 1.0 Strict"] = `"-//W3C//DTD XHTML 1.0 Strict//EN"`
	doctypes["XHTML 1.0 Transitional"] = `"-//W3C//DTD XHTML 1.0 Transitional//EN"`
	doctypes["XHTML 1.0 Frameset"] = `"-//W3C//DTD XHTML 1.0 Frameset//EN"`
	doctypes["XHTML 1.1"] = `"-//W3C//DTD XHTML 1.1//EN"`
	doctypes["HTML 5"] = `<!DOCTYPE html>`
}

func CheckDoctype(html string) string {
	var version = "UNKNOWN"

	for doctype, matcher := range doctypes {
		match := strings.Contains(html, matcher)

		if match == true {
			version = doctype
		}
	}

	return version
}
