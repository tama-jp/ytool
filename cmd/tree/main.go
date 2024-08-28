package main

import (
	"encoding/json"
	"fmt"
	"github.com/tama-jp/ytool/dir"
)

func main() {

	fmt.Println("------------ begin ------------")

	fmt.Println("------------ 1 ------------")
	//
	//treeArray := ytool.GetTree("./")
	//
	//tree, _ := json.Marshal(treeArray)
	//fmt.Println(string(tree)) // []byte型なのでstringに変換

	fmt.Println("------------ 2 ------------")

	treeArray2 := dir.Tree{Name: "./"}

	treeArray2.GetTree()

	tree2, _ := json.Marshal(treeArray2)
	fmt.Println(string(tree2)) // []byte型なのでstringに変換

	fmt.Println("------------ end ------------")

	fmt.Println("テスト")
}
