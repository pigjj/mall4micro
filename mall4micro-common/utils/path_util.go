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

var PathSplitFlag = "/"

func init() {
	// windows下goland中运行or调试模式路径获取为linux模式
	if os.Getenv("mall4micro") != "PROD" && os.Getenv("mall4micro") != "UAT" {
		return
	}
	if runtime.GOOS == "windows" {
		PathSplitFlag = "\\"
	} else {
		PathSplitFlag = "/"
	}
}

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

//
// ProjectBasePath
// @Description: 获取项目根目录
// @Document:
// @return string
//
func ProjectBasePath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		dir = getCurrentAbPathByCaller()
	}
	var pathList []string
	var basePath string
	pathList = strings.Split(dir, PathSplitFlag)
	basePath = strings.Join(projectAbPath(pathList), PathSplitFlag)
	return basePath
}

//
// IsDir
// @Description: 判断给定路径是否是文件夹
// @Document:
// @param path
// @return bool
//
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
