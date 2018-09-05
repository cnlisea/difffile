package common

import "testing"

func TestFileMD5(t *testing.T) {
	m, err := FileMD5("D:/server_yuanlai.zip")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("md5:", m)
}
