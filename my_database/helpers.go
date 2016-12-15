package my_database

import "regexp"

func isNumber(checked string) bool {
	match, _ := regexp.MatchString("[0-9]+", checked)
	return match
}