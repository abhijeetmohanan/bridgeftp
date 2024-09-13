package utils

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

func FtpParamsValidator(params map[string]string) bool {
	for _, v := range params {
		if NullChecker(v) {
			return true
		}
	}
	return false
}
