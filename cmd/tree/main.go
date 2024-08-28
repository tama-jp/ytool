package main

import (
	"encoding/json"
	"fmt"
	"github.com/tama-jp/ytool/dir"
)

func main() {

	fmt.Println("------------ begin ------------")

	fmt.Println("------------ 1 ------------")

	treeArray2 := dir.Tree{Name: "./"}

	treeArray2.GetTree()

	tree2, _ := json.Marshal(treeArray2)
	fmt.Println(string(tree2)) // []byte型なのでstringに変換

	fmt.Println("------------ 2 ------------")
	fmt.Println("------------ 3 ------------")
	fmt.Println("------------ 4 ------------")
	fmt.Println("------------ 5 ------------")
	fmt.Println("------------ 6 ------------")

	fmt.Println("------------ end ------------")

	fmt.Println("テスト")
}
