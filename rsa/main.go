package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {

	RsaGenKey(2048)
	var data = "这是一个要进行rsa解密的数据"

	encryData := RSAEncrypt([]byte(data), []byte("public.pem"))
	fmt.Println("encData", string(encryData))
	decryData := RSADecrypt(encryData, []byte("private.pem"))
	fmt.Println("decryData", string(decryData))
}

// RSAEncrypt rsa加密
// src 要加密的数据
// 公钥文件的路径
func RSAEncrypt(src, filename []byte) []byte {
	// 1. 根据文件名将文件内容从文件中读出
	file, err := os.Open(string(filename))
	if err != nil {
		return nil
	}
	// 2. 读文件
	info, _ := file.Stat()
	allText := make([]byte, info.Size())
	file.Read(allText)
	// 3. 关闭文件
	file.Close()

	// 4. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(allText)
	if block == nil {
		return nil
	}
	// 5. 解析一个DER编码的公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	// 6. 公钥加密
	result, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	return result
}

// RSADecrypt rsa加密
// src 要解密的数据
// 私钥文件的路径
func RSADecrypt(src, filename []byte) []byte {
	// 1. 根据文件名将文件内容从文件中读出
	file, err := os.Open(string(filename))
	if err != nil {
		return nil
	}
	// 2. 读文件
	info, _ := file.Stat()
	allText := make([]byte, info.Size())
	file.Read(allText)
	// 3. 关闭文件
	file.Close()
	// 4. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(allText)
	// 5. 解析一个pem格式的私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	// 6. 私钥解密
	result, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)

	return result
}
