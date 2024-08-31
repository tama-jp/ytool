package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("------------ begin ------------")

	fmt.Println("------------ 1 ------------")

	//data := dir.GetTreeText("./")
	//
	//for _, v := range data {
	//	println(v.ViewName)
	//}

	fmt.Println("------------ 2 ------------")

	//fmt.Print("Go runs on ")
	//switch os := runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	// freebsd, openbsd,
	//	// plan9, windows...
	//	fmt.Printf("%s.\n", os)
	//}

	fmt.Println("------------ 3 ------------")

	fmt.Println("This is the first line.")
	fmt.Println("This is the second line.")
	fmt.Println("This is the third line.")
	fmt.Println("Typing some text...")

	// Wait for 2 seconds to simulate typing
	time.Sleep(2 * time.Second)

	// Move cursor to row 5, column 10 (5th line, 10th column from the left)
	fmt.Print("\033[1;20H")

	// Print text at the specified location
	fmt.Println("This text is displayed at row 1, column 20.")

	// Move cursor to row 5, column 10 (5th line, 10th column from the left)
	fmt.Print("\033[5;10H")

	// Print text at the specified location
	fmt.Println("This text is displayed at row 5, column 10.")

	// Move cursor back to the original position (let's say row 8, column 1)
	fmt.Print("\033[8;1H")
	fmt.Println("This is some more text after the insert.")

	time.Sleep(2 * time.Second)

	// Move cursor to row 5, column 10 (5th line, 10th column from the left)

	//max := 10.0
	//for i := 0.0; i < max; i++ {
	//	x := i
	//	y := math.Sqrt(15.*15.0 - x*x)
	//	fmt.Printf("\033[%d;%dH", int(x), int(y))
	//	fmt.Print("*")
	//}

	fmt.Println("------------ 4 ------------")
	// 左上にボックスを表示
	fmt.Print("\033[1;1H")
	fmt.Println("+----+")
	fmt.Println("|    |")
	fmt.Println("|    |")
	fmt.Println("|    |")
	fmt.Println("+----+")

	// 右上にボックスを表示（80列目に表示したい場合）
	fmt.Print("\033[1;75H")
	fmt.Println("+----+")
	fmt.Print("\033[2;75H")
	fmt.Println("|    |")
	fmt.Print("\033[3;75H")
	fmt.Println("|    |")
	fmt.Print("\033[4;75H")
	fmt.Println("|    |")
	fmt.Print("\033[5;75H")
	fmt.Println("+----+")

	// 左下にボックスを表示（20行目に表示したい場合）
	fmt.Print("\033[20;1H")
	fmt.Println("+----+")
	fmt.Println("|    |")
	fmt.Println("|    |")
	fmt.Println("|    |")
	fmt.Println("+----+")
	fmt.Println("------------ 5 ------------")
	fmt.Println("------------ 6 ------------")

	fmt.Println("------------ end ------------")

	fmt.Println("テスト")
}
