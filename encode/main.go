package main

import (
	"fmt"
	"golang.org/x/text/encoding/charmap" // ISO-8859-1エンコーディングを提供
	"golang.org/x/text/transform"
	"strings"
)

func main() {
	// 変換する入力文字列(UTF-8)
	input := "Olá, mundo!" // ポルトガル語で「こんにちは、世界！」

	// ISO-8859-1 へのエンコード
	encoder := charmap.ISO8859_1.NewEncoder()
	encoded, _, err := transform.String(encoder, input)
	if err != nil {
		fmt.Println("エンコードエラー：", err)
		return
	}
	fmt.Println("エンコード結果(ISO-8859-1):", encoded)

	// エンコードされた文字列を再度UTF-8にデコードする
	decoder := charmap.ISO8859_1.NewDecoder()
	decoded, _, err := transform.String(decoder, encoded)
	if err != nil {
		fmt.Println("デコードエラー：", err)
		return
	}
	fmt.Println("デコード結果(UTF-8):", decoded)
}
