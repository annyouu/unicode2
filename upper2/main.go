package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"golang.org/x/text/transform"
)

// 小文字を大文字に変換するUpper型
type Upper struct{}

// Transformメソッドを実装(小文字を大文字に変換する)
func (Upper) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for i, b := range src {
		r := rune(b)
		if unicode.IsLower(r) {
			r = unicode.ToUpper(r)
		}
		dst[i] = byte(r)
		nDst++
		nSrc++
	}
	return nDst, nSrc, nil
}

// Resetメソッド
func (Upper) Reset() {}

func main() {
	// Upper型を作成
	var upper Upper

	// 変換対象の文字列をio.Readerにする
	sr := strings.NewReader("Hello, World")

	// transform.NewReaderでsrを変換しながら読み込む io.Readerを作成
	r := transform.NewReader(sr, &upper)

	// 標準出力に変換後の文字列をコピー
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Println("エラー:", err)
	}
}