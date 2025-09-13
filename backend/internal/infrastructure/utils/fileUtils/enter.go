package fileUtils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"
	"strings"
)

var (
	FILE_NAME_ERROR = errors.New("文件名错误")
	FILE_TYPE_ERROR = errors.New("文件格式错误")
)
var whiteList = []string{"jpg", "png", "jpeg", "gif", "webp"}

func FileSuffixJudge(fileName string) (string, error) {
	list := strings.Split(fileName, ".")
	length := len(list)
	if length == 1 {
		return "", FILE_NAME_ERROR
	}
	suffix := list[length-1]
	for i := range whiteList {
		if suffix == whiteList[i] {
			return suffix, nil
		}
	}
	return "", FILE_TYPE_ERROR
}

func Md5(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

func FileMd5(file string) (h string, err error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	h = Md5(data)
	return h, nil
}
