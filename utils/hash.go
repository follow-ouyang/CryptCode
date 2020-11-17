package utils

import (
	"crypto/md5"
	"crypto/sha256"
)

/*
使用md5哈希函数进行hash
 */
func Md5Hash(data []byte) []byte {
	md5Hash := md5.New()
	md5Hash.Write(data)
	hashData := md5Hash.Sum(nil)
	return hashData
}

func Sha256Hash(data []byte) []byte {
	sha256hash := sha256.New()
	sha256hash.Write(data)
	return sha256hash.Sum()
}