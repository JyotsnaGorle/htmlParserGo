package htmlVersion

import (
	"testing"
)

func TestFindHtmlVersion(t *testing.T) {
	t.Log("Running : test html version")

	htmlString := `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`

	version := CheckDoctype(htmlString)

	if version != "HTML 4.01 Strict" {
		t.Errorf("Wrong html version")
	}

}
