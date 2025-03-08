package main

import (
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
	"unicode/utf8"
)

type upperTransformer struct{}

func (t upperTransformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	// srcをルーン単位で処理
	for nSrc < len(src) {
		r, size := utf8.DecodeRune(src[nSrc:])
		if r == utf8.RuneError && size == 1 {
			// 不正な時のエラー処理
			return nDst, nSrc, transform.ErrShortSrc
		}

		// ルーンを大文字に変換
		upper := unicode.ToUpper(r)

		// 変換後dstにエンコード
		n := utf8.EncodeRune(dst[nDst:], upper)
		nDst += n
		nSrc += size
	}
	return nDst, nSrc, nil
}

func (t upperTransformer) Reset() {
	// 何もしない
}


func main() {
	t := transform.Chain(
		upperTransformer{}, //大文字変換
		norm.NFD,
	)

	input := []byte("héllo world")
	output, _, err := transform.Bytes(t, input)
	if err != nil {
		fmt.Println("変換エラー:", err)
		return
	}
	fmt.Println("変換後の文字列:", string(output))
}