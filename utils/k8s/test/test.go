package main

import (
	"fmt"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func main() {

	words := []string{"cartwheel", "foobar", "wheel", "baz"}
	match1 := fuzzy.Find("az", words) // [cartwheel wheel]

	fmt.Println("match1: ", match1)

	all := "dfafdafdfdf"
	sub := ""

	match2 := strings.Contains(all, sub)
	fmt.Println("match2: ", match2)

}
