package main

import (
	"fmt"
	"github.com/rivo/uniseg"
)

func main() {
	// 家族絵文字
	s := "👨‍👩‍👦"

	for _, c := range s {
		fmt.Printf("%x", c)
	}
	fmt.Println()

	//書記素クラスタを使って絵文字を表示
	gr := uniseg.NewGraphemes(s)
	for gr.Next() {
		fmt.Printf("%s %x\n", gr.Str(), gr.Runes())
	}
}

// func main() {
// 	gr := uniseg.NewGraphemes("Cafe\u0301")

// 	// 書記素クラスタを順に処理
// 	for gr.Next() {
// 		fmt.Printf("%s", gr.Str())

// 		for _, r := range gr.Str() {
// 			fmt.Printf("%x", r)
// 		}
// 		fmt.Println()
// 	}
// }