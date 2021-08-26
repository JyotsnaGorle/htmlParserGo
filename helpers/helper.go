package helpers

import "net/url"

/* IsValidUrl: checks if a url is valid
   Param: toTest (string) - url to check
   Returns: boolean
   sample valid url: `http://www.domain-address.com`
*/
func IsValidUrl(toTest string) bool {
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
