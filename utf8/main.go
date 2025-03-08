package main

import (
	"fmt"
	"unicode/utf8"
)

// func main() {
// 	runeValue := '世'

// 	buf := make([]byte, 3)
// 	n := utf8.EncodeRune(buf, runeValue)

// 	fmt.Printf("エンコード後: %v\n", buf)        // バイト列
// 	fmt.Printf("エンコード後文字列: %q\n", string(buf)) 
// 	fmt.Printf("エンコードされたバイト数: %d\n", n)   
// }

func main() {
	// unicodeのコーポポイント"世"をUTF-8バイト列にエンコード
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, '世')
	fmt.Printf("%v %q %d\n", buf, string(buf), n)
	// 出力：　[228 184 150] "世" 3

	// UTF-8エンコードされた文字列"世界"のデコードする
	b := []byte("世界")
	// 世と界のそれぞれのデコードを行う
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%q:%v", r, size)
		b = b[size:] // 残りのバイト列に進む
	}
}