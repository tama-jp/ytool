package dir

import (
	"os"
	"path/filepath"
)

type Tree struct {
	Name     string `json:"name"`      // ファイル・ディレクトリ名
	FullPath string `json:"full_path"` // フルパス
	DirCheck bool   `json:"dir_check`  // ディレクトリチェック
	Next     []Tree `json:"next"`      // 配下のパス
}

func (path *Tree) GetTree() Tree {
	var paths []Tree

	dirStr := path.Name

	tree := getTreeSub(dirStr, paths)

	dirStr2, err := filepath.Abs(dirStr)

	if err != nil {
		panic(err)
	}

	file, _ := os.Stat(dirStr2)

	dirCheck := false

	if file.IsDir() {
		dirCheck = true
	}

	*path = Tree{
		Name:     dirStr,
		FullPath: dirStr2,
		DirCheck: dirCheck,
		Next:     tree,
	}

	return *path
}

func getTreeSub(dirStr string, treeArray []Tree) []Tree {
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

		if file.IsDir() {
			paths = filepath.Join(dirStr, file.Name())
			next = getTreeSub(paths, next)
			dirCheck = true
		}

		dataTemp := Tree{
			Name:     file.Name(),
			FullPath: filepath.Join(dirStr, file.Name()),
			DirCheck: dirCheck,
			Next:     next,
		}

		treeArray = append(treeArray, dataTemp)
	}

	return treeArray
}
