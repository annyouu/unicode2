package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/width"
)

// Shift_JISのファイルの全角英数を半角に、半角カナを全角に変換する
func foldShiftJISFile(filename string) error {
	// ファイルを開く
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Shift_JSIからUTF-8に変換しつつ、半角、全角を統一する
	dec := japanese.ShiftJIS.NewDecoder() // Shift_JISからUTF-8へ
	t := transform.Chain(dec, width.Fold) // foldで全角を半角に、半角を全角に変換
	r := transform.NewReader(f, t)

	// 変換された内容を1行ずつ読み込んで表示
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Println(s.Text())  // 変換後のテキストを出力
	}

	// 読み込んだ時エラーだったら、エラーを出す
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	// コマンドラインでファイル名を取得
	if len(os.Args) < 2 {
		fmt.Println("go run main.go <Shift_JISのテキストファイル>")
		os.Exit(1)
	}
	filename := os.Args[1]
	
	if err := foldShiftJISFile(filename); err != nil {
		fmt.Println("エラー:", err)
		os.Exit(1)
	}
}