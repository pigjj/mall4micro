package utils

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

//
// projectAbPath
// @Description: 通过给定路径截取所需项目名称的绝对路径
// @Document:
// @param pathList
// @return []string
//
func projectAbPath(pathList []string) []string {
	for i, v := range pathList {
		if v == "mall4micro" {
			return pathList[:i+1]
		}
	}
	panic(errors.New("str 'mall4micro' not found"))
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func ProjectBasePath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		dir = getCurrentAbPathByCaller()
	}
	var pathList []string
	var basePath string
	if runtime.GOOS == "windows" {
		pathList = strings.Split(dir, "\\")
		basePath = strings.Join(projectAbPath(pathList), "\\")
	} else {
		pathList = strings.Split(dir, "/")
		basePath = strings.Join(projectAbPath(pathList), "/")
	}
	return basePath
}
