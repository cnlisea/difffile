package export

import (
	"difffile/internal/common"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExportFile(paths []string, basePath string, exportPath string) error {
	var (
		path string
		err  error
	)
	for i := range paths {
		path = strings.Replace(paths[i], basePath+common.PathSep, "", 1)
		path = exportPath + common.PathSep + path
		// 获取文件所属目录
		dir := filepath.Dir(path)
		// 递归创建目录
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		if _, err = common.CopyFile(paths[i], path); err != nil {
			return err
		}

		fmt.Println("export file:", paths[i], "to", path, "success!!!")
	}
	return nil
}
