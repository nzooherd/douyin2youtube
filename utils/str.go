package utils

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"strings"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}


func HashCode(s string) int{
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func Mp4ToTsFileName(name string) string{
	return strings.Replace(name, ".mp4", ".ts", -1)
}
