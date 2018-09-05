package export

import (
	"os"
	"path/filepath"
)

func ExportFile(paths []string) error {
	var err error
	for i := range paths {
		// 获取文件所属目录
		dir := filepath.Dir(paths[i])
		// 递归创建目录
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}

	}
	return nil
}
