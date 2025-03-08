package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/utf8"
	"io"
)

func main() {
	// 変換用のインタフェースを取得（UTF-8 -> ISO-8859-1）
	utf8ToISO := transform.Chain(utf8.NewDecoder(), transform.RemoveFunc(func(r rune) rune {
		if r == 0xFFFD { // 置換文字（U+FFFD）を削除
			return -1
		}
		return r
	}))

	// 変換元のデータ（UTF-8）
	src := []byte("こんにちは")

	// 出力バッファを作成
	var dst bytes.Buffer

	// 変換処理
	_, _, err := utf8ToISO.Transform(&dst, src, true)
	if err != nil {
		fmt.Println("変換エラー:", err)
		return
	}

	// 変換後の結果を表示
	fmt.Println("変換後:", dst.String())
}
