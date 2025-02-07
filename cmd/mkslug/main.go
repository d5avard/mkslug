package main

import (
	"fmt"
	"os"

	"github.com/d5avard/mkslug/internal/slug"
)

func main() {
	args := os.Args[1:]
	var words []string
	var result string

	words = slug.CleanWords(args)
	result = slug.MakeSlug(words)

	fmt.Println(result)
}
