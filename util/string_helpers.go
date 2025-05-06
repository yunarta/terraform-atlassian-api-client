package util

import "fmt"
import "net/url"

func QueryParam(key, value string) string {
	if len(value) > 0 {
		return fmt.Sprintf("&%s=%s", key, url.QueryEscape(value))
	} else {
		return ""
	}
}

func CoalesceString(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
