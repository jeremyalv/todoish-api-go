package constants

import "regexp"

var (
	RegexTodoId = regexp.MustCompile("/^[0-9A-F]{8}-[0-9A-F]{4}-1[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i")
)
