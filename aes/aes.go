package aes

import (
	"crypto/aes"
	"CryptCode/utils"
	"crypto/cipher"
	"fmt"
)

/**
 * 使用AES算法对明文进行加密
 */
func AESEnCrypt(origin []byte, key []byte) ([]byte, error) {
	//3元素：key、data、mode
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//2、对明文数据进行尾部填充
	cryptData := utils.PKCS5EndPadding(origin, block.BlockSize())
	//3、实例化加密mode
	blockMode := cipher.NewCBCEncrypter(block, key)
	//4、加密
	cipherData := make([]byte, len(cryptData))
	blockMode.CryptBlocks(cipherData, cryptData)
	return cipherData, nil
}

/*
该函数用于AES解密
 */
func AESDnCrypt(origin []byte,key []byte) ([]byte,error) {
	//三要素：key、data、mode
	block,err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	blockMode := cipher.NewCBCDecrypter(block,key)
	cipherData := make([]byte,len(origin))
	blockMode.CryptBlocks(cipherData,origin)
	cipherData = utils.ClearPKCS5Padding(cipherData,block.BlockSize())
	return cipherData,nil
}