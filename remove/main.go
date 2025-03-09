package main

import (
	"fmt"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)


func main() {
	// カテゴリに属する文字を削除
	removeNumbers := runes.Remove(runes.In(unicode.Nd))

	// 変換
	s, _, err := transform.String(removeNumbers, "Go123Lang456")
	if err != nil {
		fmt.Println("エラ-:", err)
		return
	}
	// 出力
	fmt.Println(s)
}

// func main() {
// 	// Unicodeの削除
// 	removeMn := runes.Remove(runes.In(unicode.Mn))

// 	// 正規化(分解、統合文字削除、正規化)
// 	t := transform.Chain(norm.NFD, removeMn, norm.NFC)

// 	// "résumé"という文字列を変換
// 	s , _, err := transform.String(t, "résumé")
// 	if err != nil {
// 		fmt.Println("エラー:", err)
// 		return
// 	}

// 	// 結果の出力
// 	fmt.Println(s)
// }