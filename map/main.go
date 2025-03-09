package main

import (
	"fmt"
	
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

func main() {
	// 小文字を大文字に変換するTransformerを作成
	t := runes.Map(func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return r - ('a' - 'A')
		}
		return r // それ以外は変換しない
	})

	// 変換
	s, _, err := transform.String(t, "Hello, World")
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// 出力
	fmt.Println(s) // "HELLO, WORLD"
}