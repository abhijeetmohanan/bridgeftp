package utils

import (
	"log"
	"net/url"
)

func ParseInput(key string, furl string) (map[string]string, error) {
	// Define return variable

	params := make(map[string]string)
	//

	// Parse the URLs
	sourceURL, err := url.Parse(furl)
	if err != nil {
		return params, err
	}

	// Output the parsed URLs

	// Extract and print the protocol (scheme)
	params["scheme"] = sourceURL.Scheme
	params["username"] = sourceURL.User.Username()
	sourcePassword, _ := sourceURL.User.Password()
	params["password"] = sourcePassword
	params["host"] = sourceURL.Host
	params["path"] = sourceURL.Path

	// log the paramters used for connection
	// ToDo: sorted map implementation
	parameters := ""

	for k, v := range params {
		parameters = parameters + k + ": " + v + ", "
	}

	log.Println(key + ":: " + parameters)

	return params, nil
}

func NullChecker(key string, ftpurl string) {
	// Check if string is null
	if ftpurl == "" {
		log.Panicf("Error: Empty URL for %s", key)
	}
}

func SchemeValidator(scheme string) bool {
	log.Printf("scheme: %s", scheme)
	if scheme == "ftp" {
		return true
	}
	return false
}

func FtpParamsValidator(params map[string]string) {
	for k, v := range params {
		if v == "" {
			log.Panicf("Null Values passed for %v", k)
		}
	}
}
