package slug

import (
	"strings"
	"unicode"
)

func CleanWords(words []string) []string {
	ls := make([]string, 0, len(words))
	for _, word := range words {
		r := strings.Map(func(rn rune) rune {
			switch {
			case unicode.IsSymbol(rn) || unicode.IsPunct(rn):
				return '-'
			case unicode.IsLetter(rn) || unicode.IsNumber(rn):
				return rn
			default:
				return -1
			}
		}, word)
		ls = append(ls, r)
	}
	return ls
}

func MakeSlug(words []string) string {
	var r string
	r = strings.Join(words, "-")
	r = strings.ToLower(r)
	return r
}
