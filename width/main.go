package main

import (
	"fmt"
	"golang.org/x/text/width"
)

func main() {
	// 対象とする文字のスライス
	rs := []rune{'５', 'ｱ', 'ア', 'A', 'α'}

	// ヘッダーを表示
	fmt.Println("rune\tWide\tNarrow\tFolded\tKind")
	fmt.Println("--------------------------------------------------")

	// 各文字について幅情報を表示
	for _, r := range rs {
		p := width.LookupRune(r) // width.Properties型の値を取得
		w, n, f, k := p.Wide(), p.Narrow(), p.Folded(), p.Kind() //幅情報を取得

		// 幅情報を整形して表示
		fmt.Printf("%2c\t%2c\t%3c\t%3c\t%s\n", r, w, n, f, k)
	}
}