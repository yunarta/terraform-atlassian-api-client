package util

import "fmt"

func QueryParam(key, value string) string {
	if len(value) > 0 {
		return fmt.Sprintf("&%s=%s", key, value)
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
