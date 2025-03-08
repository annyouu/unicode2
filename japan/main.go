package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"io"
	"os"
)

func printCSV(filename string) error {
	// ファイルを開く
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Shift_JISとして読み込むデコーダを作成
	dec := japanese.ShiftJIS.NewDecoder()

	// デコーダを使ってCSVリーダーを作成
	cr := csv.NewReader(dec.Reader(f))

	// CSVを1行ずつ読み込む
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(rec)
	}
	return nil
}


func main() {
	// 例として"shiftjis.csvというSHift_JISエンコードのCSVファイルを読み込む"
	err := printCSV("shiftjis.csv")
	if err != nil {
		fmt.Println("エラー:", err)
	}
}