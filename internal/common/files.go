package common

import (
	"io/ioutil"
	"os"
	"strings"
)

// 平台分隔符号
const PathSep = string(os.PathSeparator)

func GetDirFiles(path string, fileSuffix ...string) ([]string, error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, fi := range dir {
		// 目录, 递归遍历
		if fi.IsDir() {
			f, err := GetDirFiles(path + PathSep + fi.Name())
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
				continue
			}
		}

		// 无过滤格式
		if len(fileSuffix) == 0 {
			files = append(files, path+PathSep+fi.Name())
		}
	}

	return files, nil
}
