package helpers

import "strings"

func split(s string) []string {
	return strings.Split(s, "/n")
}
