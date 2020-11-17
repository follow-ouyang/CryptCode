package main

import (
	"CryptCode/3des"
	"CryptCode/aes"
	"CryptCode/des"
	"CryptCode/rsa"
	"fmt"
)

func main() {
	//des: 块加密
	//des: key、data、mode
	/**
	 * key：密钥
	 * data：明文，即将加密的明文
	 * mode：两种模式，加密和解密
	 */
	//key := []byte("c1906041")
	//
	//data := "遇贵人先立业，遇良人先成家，遇阿姨成家立业!"
	//
	//////加密：crypto
	//block, err := des.NewCipher(key)
	//
	//if err != nil {
	//	panic("初始化密钥错误，请重试！")
	//}
	////dst, src
	//dst := make([]byte, len([]byte(data)))
	////加密过程
	//block.Encrypt(dst, []byte(data))
	//
	//fmt.Println("加密后的内容：", dst)
	//
	////解密
	//deBlock, err := des.NewCipher(key)
	//if err != nil {
	//	panic("初始化密钥错误，请重试！")
	//}
	//deData := make([]byte, len(dst))
	//deBlock.Decrypt(deData, dst)
	//
	//fmt.Println(string(deData))

	//一、对数据进行加密  DES秘钥为8字节、3DES秘钥长度为24字节
	//key := []byte("c1906041")
	//
	//data := "遇贵人先立业，遇良人先成家"

	//1、得到cipher
	//block, err := des.NewCipher(key)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	////2、对数据明文进行结尾块填充
	//paddingData := PKCS5Padding([]byte(data), block.BlockSize())
	////3、选择模式
	//mode := cipher.NewCBCEncrypter(block, key)
	////4、加密
	//dstData := make([]byte, len(paddingData))
	//mode.CryptBlocks(dstData, paddingData)
	//
	//fmt.Println("加密后的内容：", string(dstData))

	//二、对数据进行解密
	//DES三元素：key、data、mode
	//key1 := []byte("c1906041")
	//data1 := dstData //待解密的数据
	//block1, err := des.NewCipher(key1)
	//if err != nil {
	//	panic(err.Error())
	//}
	//deMode := cipher.NewCBCDecrypter(block1, key1)
	//originalData := make([]byte, len(data1))
	////分组解密
	//deMode.CryptBlocks(originalData, data1)
	//originalData = utils.ClearPKCS5Padding(originalData, block1.BlockSize())
	//fmt.Println("解密后的内容:", string(originalData))

	//一、使用DES进行加解密
	fmt.Println("DES算法")
	key := []byte("20201112") //des秘钥长度：8字节
	data := "窗含西岭千秋雪，门泊东吴万里船"
	cipherText, err := des.DESEnCrypt([]byte(data), key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	originText, err := des.DESDeCrypt(cipherText, key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("DES解密后：", string(originText))

	//二、3DES加解密
	fmt.Println("3DES算法")
	key1 := []byte("202011122020111220201112") //3des的秘钥长度是24字节
	data1 := "穷则独善其身，富则兼济天下"
	//加密
	cipherText1, err := _des.TripleDESEncrypt([]byte(data1), key1)
	if err != nil {
		fmt.Println("3DES加密失败:", err.Error())
		return
	}
	//解密
	originalText1, err := _des.TripleDESDecrypt(cipherText1, key1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("3DES解密后的内容：", string(originalText1))

	//三、AES算法:
	fmt.Println("AES算法")
	key2 := []byte("2020111220201112") //8
	data2 := "天高皇帝远"
	//加密
	cipherText2, err := aes.AESEnCrypt([]byte(data2), key2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//解密
	originText3,err := aes.AESDnCrypt(cipherText2,key2)
	if err != nil {
		fmt.Println("aes解密错误",err.Error())
		return
	}
	fmt.Println("aes解密后的内容：",string(originText3))

	//四、RSA算法的使用
	fmt.Println("RSA算法")
	data4 := "把我的悲伤留给自己"
	//4.1  生成密钥对
	priv,err := rsa.CreatePairKeys()
	if err != nil {
		fmt.Println("生成密钥对出错：",err.Error())
		return
	}
	//4.2  对数据进行加密
	cipherText4,err := rsa.RSAEncrypt(priv.PublicKey,[]byte(data4))
	if err != nil {
		fmt.Println("rsa算法加密失败：",err.Error())
		return
	}
	originalText4,err := rsa.RSADecrypt(priv,cipherText4)
	if err != nil {
		fmt.Println("rsa算法解密失败：",err.Error())
		return
	}
	fmt.Println("rsa解密成功，结果是：",string(originalText4))

	//4.3  对原文数据进行签名
	signText4,err := rsa.RSASign(priv,[]byte(data4))
	if err != nil {
		fmt.Println("rsa算法签名失败：",err.Error())
		return
	}
	//4.4  对签名数据进行验证
	verfityResult,err := rsa.RSAVerify(priv.PublicKey,[]byte(data4),signText4)
	if err != nil {
		fmt.Println("rsa算法签名验证失败：",err.Error())
		return
	}
	if verfityResult {
		fmt.Println("rsa签名验证成功")
	}else {
		fmt.Println("rsa签名验证成功")
	}

}

/**
 * 明文数据尾部填充
 */
//func PKCS5Padding(text []byte, blockSize int) []byte {
//	//计算要填充的快内容的大小
//	paddingSize := blockSize - len(text)%blockSize
//	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
//	//fmt.Println("明文尾部的内容：",paddingText)
//	return append(text, paddingText...)
//}
