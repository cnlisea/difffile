package diff

import (
	"difffile/internal/common"
	"difffile/internal/initialize"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Diff struct {
	readPath   string   // 指定目录位置
	fileSuffix []string // 过滤文件后缀
}

func NewDiff(readPath string, fileSuffix ...string) *Diff {
	return &Diff{readPath: readPath, fileSuffix: fileSuffix}
}

func (d *Diff) Diff() ([]string, error) {
	var fileNotExist bool
	// 读取原文件
	fileData, err := ioutil.ReadFile(d.readPath + common.PathSep + initialize.DefaultWriteFileName)
	switch err {
	case nil:
	case os.ErrNotExist:
		fileNotExist = true
	default:
		return nil, err
	}

	oldDataMap := make(map[string]string)
	if !fileNotExist {
		// 序列化
		if err = json.Unmarshal(fileData, &oldDataMap); err != nil {
			return nil, err
		}
	}

	// 获取当前所有文件
	files, err := common.GetDirFiles(d.readPath, d.fileSuffix...)
	if err != nil {
		return nil, err
	}

	// md5文件做对比
	var (
		diffMap    = make([]string, 0)
		newM, oldM string
		ok         bool
	)
	for i := range files {
		if files[i] == initialize.DefaultWriteFileName {
			continue
		}
		newM, err = common.FileMD5(files[i])
		if err != nil {
			return nil, err
		}

		if oldM, ok = oldDataMap[files[i]]; !ok || oldM != newM {
			diffMap = append(diffMap, files[i])
		}
	}

	return diffMap, nil
}
