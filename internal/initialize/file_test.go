package initialize

import "testing"

func TestNewInit(t *testing.T) {
	init := NewInit("D:/go_work/src/test", "D:/go_work/src/test", "xml")
	if err := init.Exec(); err != nil {
		t.Fatal(err)
	}
}
