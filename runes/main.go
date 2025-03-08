package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"unicode"
)

func main() {
	// 変換対象の文字列(カタカナとその他の文字を含む)
	text := "こんにちは ゴハン デス！"

	// カタカナの文字だけを大文字化するTransformer
	katakataTransformer := runes.If(
		runes.In(unicode.Katakana),
		runes.Map(unicode.ToUpper),
		transform.Nop,
	)

	// 変換の適用適用
}