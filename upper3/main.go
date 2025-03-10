package main

import (
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"os"
	"strings"
)

// 小文字を大文字に変換する
type Upper struct{}

func (Upper) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for {
		// srcを全て処理し終えた、またはdstが足りなくなった時
		if len(src) <= nSrc || len(dst) <= nDst {
			return
		}

		// 小文字から大文字へ変換する
		if 'a' <= src[nSrc] && src[nSrc] <= 'z' {
			dst[nDst] = src[nSrc] - ('a' - 'A')
		} else {
			dst[nDst] = src[nSrc]
		}
		// 更新
		nSrc++
		nDst++
	}
}

func (Upper) Reset() {}

func main() {
	// 変換対象の文字列
	sr := strings.NewReader("hello, world!")

	// transform.NewReaderで大文字変換しながら読み込む
	r := transform.NewReader(sr, Upper{})

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Println("エラー:", err)
	}
}
