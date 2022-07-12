package autils

import "strings"

func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func IsStringContainsSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.Contains(a, b) {
			return true
		}
	}

	return false
}

func StringPad(val string, direction string, length int) string {
	res := val
	addition := ""

	for i := 2; i < length; i++ {
		if len(val) < i {
			for j := 0; j < length-i; j++ {
				addition += "0"
			}

			if direction == "left" {
				res = addition + val
			} else {
				res = val + addition
			}

			break
		}
	}

	return res
}

func Contains(s []string, search string) bool {
	for _, v := range s {
		if v == search {
			return true
		}
	}

	return false
}
