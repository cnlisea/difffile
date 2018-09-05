package common

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
)

//EncryptMD5 encrypt given []byte with MD5 algorithm
func FileMD5(filePath string) (string, error) {
	if filePath == "" {
		return "", errors.New("No such file or directory")
	}

	// read file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil)), nil
}
