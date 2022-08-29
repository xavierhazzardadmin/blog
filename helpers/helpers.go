package helpers

import "strings"

func SplitRows(s string) []string {
	return strings.Split(s, "\n\n")
}
