package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Tree struct {
	Name      string `json:"name"`      // ファイル・ディレクトリ名
	FullPath  string `json:"full_path"` // フルパス
	DirCheck  bool   `json:"dir_check"` // ディレクトリチェック
	Hierarchy int    `json:"hierarchy"` // ディレクトリ階層
	Next      []Tree `json:"next"`      // 配下のパス
}

type TextTree struct {
	ViewName string // ファイル・ディレクトリ名
	FullPath string // フルパス
}

func GetTreeText(pathStr string) []TextTree {
	path := GetTreeNode(pathStr)

	var nextDir []bool

	var output []TextTree
	data := makeTextTree(path.Next, nextDir, output)

	return data
}

func makeTextTree(input []Tree, nextDir []bool, output []TextTree) []TextTree {

	for count := 0; count < len(input); count++ {
		data := input[count]

		var textTemp = ""

		for _, v := range nextDir {
			if v == true {
				textTemp += "│   "
			} else {
				textTemp += "    "
			}
		}

		if count == len(input)-1 {
			textTemp += fmt.Sprintf("└── %s", data.Name)
		} else {
			textTemp += fmt.Sprintf("├── %s", data.Name)
		}

		var textTreeTemp = TextTree{
			ViewName: textTemp,
			FullPath: data.FullPath,
		}

		output = append(output, textTreeTemp)

		if data.Next != nil {
			nextDirTemp := nextDir
			if count == len(input)-1 {
				nextDirTemp = append(nextDirTemp, false)
			} else {
				nextDirTemp = append(nextDirTemp, true)
			}
			output = makeTextTree(data.Next, nextDirTemp, output)
		}

	}
	return output
}

func GetTreeNode(path string) Tree {
	var paths []Tree

	dirStr := path

	hierarchy := 0
	tree := getTreeSub(dirStr, hierarchy+1, paths)

	dirStr2, err := filepath.Abs(dirStr)

	if err != nil {
		panic(err)
	}

	file, _ := os.Stat(dirStr2)

	dirCheck := false

	if file.IsDir() {
		dirCheck = true
	}

	return Tree{
		Name:      dirStr,
		FullPath:  dirStr2,
		DirCheck:  dirCheck,
		Hierarchy: hierarchy,
		Next:      tree,
	}
}

func getTreeSub(dirStr string, hierarchy int, treeArray []Tree) []Tree {
	var err error
	dirStr, err = filepath.Abs(dirStr)

	if err != nil {
		panic(err)
	}

	files, err := os.ReadDir(dirStr)
	if err != nil {
		panic(err)
	}

	var paths string

	for _, file := range files {
		dirCheck := false
		var next []Tree

		str := file.Name()

		if len(str) <= 0 {
			continue
		}

		if strings.Compare(".", string([]rune(str)[:1])) == 0 {
			continue
		}

		if file.IsDir() {
			paths = filepath.Join(dirStr, file.Name())
			next = getTreeSub(paths, hierarchy+1, next)
			dirCheck = true
		}

		dataTemp := Tree{
			Name:      file.Name(),
			FullPath:  filepath.Join(dirStr, file.Name()),
			DirCheck:  dirCheck,
			Hierarchy: hierarchy,
			Next:      next,
		}

		treeArray = append(treeArray, dataTemp)
	}

	return treeArray
}
