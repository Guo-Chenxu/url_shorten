package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func StrToMd5(str string) string {
	hasher := md5.New()
	_, _ = io.WriteString(hasher, str)

	hashBytes := hasher.Sum(nil)
	hashStr := fmt.Sprintf("%x", hashBytes)
	return hashStr
}
