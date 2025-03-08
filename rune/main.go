package main

import "fmt"

func main() {
	str := "こんにちは"
	runes := []rune(str)

	// 文字列の各文字を表示
	for i, r := range runes {
		fmt.Printf("Index: %d, Rune: %c, Code Point: %U\n", i, r, r)
	}
}