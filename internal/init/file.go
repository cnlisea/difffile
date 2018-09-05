package init

import (
	"difffile/internal/common"
	"encoding/json"
	"os"
)

// 默认保存的md5文件
const DefaultWriteFileName = ".difffile.json"

type Init struct {
	readPath  string
	writePath string

	fileSuffix []string // 过滤文件
}

func NewInit(readPath string, writePath string, fileSuffix ...string) *Init {
	return &Init{readPath: readPath, writePath: writePath, fileSuffix: fileSuffix}
}

func (i *Init) Exec() error {
	// 获取所有文件
	files, err := common.GetDirFiles(i.readPath, i.fileSuffix...)
	if err != nil {
		return err
	}

	var (
		md5Map = make(map[string]string, len(files))
		m      string
	)
	// md5文件
	for i := range files {
		if files[i] == DefaultWriteFileName {
			continue
		}
		m, err = common.FileMD5(files[i])
		if err != nil {
			return err
		}
		md5Map[files[i]] = m
	}

	// 保存json文件 以读写方式打开文件，如果不存在，则创建
	writeFile, err := os.OpenFile(i.writePath+common.PathSep+DefaultWriteFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		return err
	}
	defer writeFile.Close()

	// 将所有文件md5值保存至文件
	encode := json.NewEncoder(writeFile)
	encode.SetIndent("", "  ")
	if err = encode.Encode(md5Map); err != nil {
		return nil
	}

	return nil
}
