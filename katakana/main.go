package main

import (
	"fmt"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
	"unicode"
)

func main() {
	t := runes.If(
		runes.In(unicode.Katakana), // カタカナを対象
		width.Widen,  // 全角に変換
		nil,
	)

	// 変換対象の文字列
	text := "５ｱアAα"

	// 変換の適用
	result, _, _ := transform.String(t, text)

	// 出力
	fmt.Println("元の文字列:", text)
	fmt.Println("変換後:", result)
}