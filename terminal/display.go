package terminal

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// ConsoleScreenBufferInfo は、Windowsコンソールバッファの情報を格納します
type ConsoleScreenBufferInfo struct {
	Size           int16
	CursorPosition int16
	Attributes     int16
	Window         struct {
		Left   int16
		Top    int16
		Right  int16
		Bottom int16
	}
	MaximumWindowSize int16
}

func getTerminalSize() (int, int, error) {
	os := runtime.GOOS
	fmt.Printf("Operating System: %s\n", os)

	// OSごとの処理
	switch os {
	case "windows":
		//fmt.Println("You are using Windows.")
		return getTerminalSizeForForWindows()
	case "darwin":
		//fmt.Println("You are using macOS.")
		return getTerminalSizeForUnixMac()
	case "linux":
		//fmt.Println("You are using Linux.")
		return getTerminalSizeForUnixMac()
	default:
		fmt.Println("Unknown OS.")
	}

	return 0, 0, errors.New("エラー")
}

// getTerminalSize は、ターミナルの幅と高さを返します
func getTerminalSizeForForWindows() (int, int, error) {
	// 未使用
	return 0, 0, nil
}

func getTerminalSizeForUnixMac() (int, int, error) {

	cmd := exec.Command("stty", "size")
	cmd.Stdin = nil
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	sizes := strings.Split(strings.TrimSpace(string(out)), " ")
	rows, err := strconv.Atoi(sizes[0])
	if err != nil {
		return 0, 0, err
	}
	cols, err := strconv.Atoi(sizes[1])
	if err != nil {
		return 0, 0, err
	}

	return rows, cols, nil
}
