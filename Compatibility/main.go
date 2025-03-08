package main

import (
	"fmt"
	"golang.org/x/text/unicode/norm"
)

func main() {
	s := "ゴ"
	fmt.Printf("%[1]q %+[1]q\n", s)

	// 互換等価性に基づいて分解（NFKD）
	s = norm.NFKD.String(s)
	fmt.Printf("%[1]q %+[1]q\n", s)
	// 互換等価性に基づいて合成（NFKC）
	s = norm.NFKC.String(s)
	fmt.Printf("%[1]q %[1]q\n", s)
}