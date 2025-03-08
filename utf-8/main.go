package main

import (
	"fmt"
	"strings"
	"io"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)


// 書き込み
func main() {
	// ファイルを作成
	file, err := os.Create("output_sjis.txt")
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	defer file.Close()

	// UTF-8 Shift_JISに変換するWriter
	tw := transform.NewWriter(file, japanese.ShiftJIS.Newcoder())

	// ファイルにShift_JISで書き込み
	_, err := tw.Write([]byte("こんにちは、 世界\n"))
	if err != nil {
		fmt.Println("エラー:", err)
	}
	fmt.Println("Shift_JIS形式でoutput_sjis.txtに保存しました")
}


//読み込み
// func main() {
// 	// UTF-8の文字列をio.Readerとして作成
// 	r := strings.NewReader("こんにちは、世界")

// 	// UTF-8 shift-jisに変換するTransformer
// 	sjisTransformer := japanese.ShiftJIS.NewEncoder()

// 	// 変換付きReaderを作成
// 	tr := transform.NewReader(r, sjisTransformer)

// 	// 変換後のデータを標準出力にコピー
// 	_, err := io.Copy(os.Stdout, tr)
// 	if err != nil {
// 		fmt.Println("エラー：", err)
// 	}
// }