package common

import (
	"errors"
	"net/http"
	"net/url"
)

func isUrl(URL string) bool {
	u, err := url.Parse(URL)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func CheckForURL(URL string) (bool, error) {
	//Leverage Go's HTTP Post function to make request
	if !isUrl(URL) {
		return false, errors.New("not in standard URL format")
	}

	_, err := http.Get(URL)

	if err != nil {
		return false, err
	}

	return true, nil
}
