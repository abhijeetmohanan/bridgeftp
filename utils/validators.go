package utils

import (
	"log"
)

func NullChecker(ftpurl string) bool {
	// Check if string is null
	return ftpurl == ""
}

func SchemeValidator(source, dest, key string) bool {
	if source == dest {
		if source == key {
			return true
		}
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
