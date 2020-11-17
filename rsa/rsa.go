package rsa

import (
	"CryptCode/utils"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"flag"
	"fmt"
)

/*
私钥：
公钥：
 */

func CreatePairKeys() (*rsa.PrivateKey,error) {
	//1、先生成私钥
	var bits int
	flag.IntVar(&bits,"b",2048,"密钥长度")

	privateKey,err := rsa.GenerateKey(rand.Reader,bits)
	if err != nil {
		return nil,err
	}

	//2、根据私钥生成公钥
	//publicKey := privateKey.Public()
	//3、讲私钥和公钥进行返回
	return privateKey,nil

}

//=======================第一种组合：公钥加密。私钥解密======================

/*
使用RSA算法对数据进行加密，返回加密后的密文
*/
func RSAEncrypt(key rsa.PublicKey,data []byte) ([]byte,error) {
	rsa
	return rsa.EncryptPKCS1v15(rand.Reader,&key,data)
}

/*
使用RSA算法对密文数据进行解密，返回解密后的铭文
 */
func RSADecrypt(private *rsa.PrivateKey,cipher []byte) ([]byte,error) {
	return rsa.DecryptPKCS1v15(rand.Reader,private,cipher)
}

//==========================第二种组合：私钥签名，公钥验签=======================

/*
使用RSA算法对数据进行数字签名，并返回签名信息
 */
func RSASign(private *rsa.PrivateKey,data []byte,hash crypto.Hash) ([]byte,error) {
	if hash == crypto.MD5 {
		hashed := utils.Md5Hash(data)
		return rsa.SignPKCS1v15(rand.Reader,private,crypto.MD5,hashed)
	}else if hash == crypto.SHA256 {
		hashed := utils.Sha256Hash(data)
		return rsa.SignPKCS1v15(rand.Reader,private,crypto.MD5,hashed)
	}else {
		fmt.Println("抱歉，暂不支持此方法")
	}

}

/*
使用RSA算法对数据进行签名验证，并返回签名验证的结果
如果验证通过，返回true
验证不通过，返回flase，同时error中有错误信息
 */
func RSAVerify(pub rsa.PublicKey,data []byte,signText []byte) (bool,error) {
	hashd := utils.Md5Hash(data)
	err := rsa.VerifyPKCS1v15(&pub,crypto.MD5,hashd,signText)
	return err==nil,err
}