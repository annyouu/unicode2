package main

import (
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
)

func main() {
	str := "５ｱアAα"

	fmt.Println("=== Folded (全角、半角で統一する) ===")
	folded, _, _ := transform.String(width.Fold, str)
	for _, r := range folded {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}

	fmt.Println("\n=== Narrow (半角に統一) ===")
	narrowed, _, _ := transform.String(width.Narrow, str)
	for _, r := range narrowed {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}

	fmt.Println("\n=== Widen (全角に統一) ===")
	widened, _, _ := transform.String(width.Widen, str)
	for _, r := range widened {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}
}