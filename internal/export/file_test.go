package export

import (
	"difffile/internal/common"
	"testing"
)

func TestExportFile(t *testing.T) {
	paths, err := common.GetDirFiles("D:/go_work/src/test")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("paths:", paths)

	if err = ExportFile(paths, "D:/go_work/src/test", "D:/go_work/src/11111111"); err != nil {
		t.Fatal(err)
	}
}
