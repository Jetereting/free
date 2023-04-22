package utils

import (
	"encoding/base64"
	"os"
)

func DecodeBase64(str string) (res string) {
	s, _ := base64.StdEncoding.DecodeString(str)
	res = string(s)
	return
}

func PathIsExist(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}
