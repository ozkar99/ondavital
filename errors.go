package ondavital

import "strings"

type ovError struct {
	message string
	elem    string
	regexp  []string
}

func (he ovError) Error() string {
	return "Message: " + he.message + "\n" +
		"Element: " + he.elem + "\n" +
		"RegeExp: " + strings.Join(he.regexp, "\n")
}
