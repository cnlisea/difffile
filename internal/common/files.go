package common

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 平台分隔符号
const PathSep = "/"

//const PathSep = string(os.PathSeparator)

func GetDirFiles(path string, fileSuffix ...string) ([]string, error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, fi := range dir {
		// 目录, 递归遍历
		if fi.IsDir() {

			f, err := GetDirFiles(path+PathSep+fi.Name(), fileSuffix...)
			if err != nil {
				return nil, err
			}
			files = append(files, f...)
			continue
		}

		// 过滤指定格式
		for j := range fileSuffix {
			if strings.HasSuffix(fi.Name(), fileSuffix[j]) {
				files = append(files, path+PathSep+fi.Name())
				break
			}
		}

		// 无过滤格式
		if len(fileSuffix) == 0 {
			files = append(files, path+PathSep+fi.Name())
		}
	}

	return files, nil
}

func CopyFile(src string, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}
