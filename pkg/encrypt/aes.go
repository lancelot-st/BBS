package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var Key = []byte("#HvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR")

////使用PKCS7进行填充
//
//func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	//将[]byte{byte(padding)}复制padding个合成新的切片
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(ciphertext, padtext...)
//}
//
//func PKCS7UnPadding(Data []byte) ([]byte, error) {
//	length := len(Data)
//	if length == 0 {
//		return nil, errors.New("加密字符串错误")
//	}
//	unPadding := int(Data[length-1])
//	return Data[:(length - unPadding)], nil
//}
//
////aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
//
//func AesCBCEncrypt(rawData, key []byte) ([]byte, error) {
//	//创建加密实列
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	//填充原文
//	blockSize := block.BlockSize()
//	rawData = PKCS7Padding(rawData, blockSize)
//	//block 大小得和iv向量一致
//	mode := cipher.NewCBCDecrypter(block, key[:blockSize])
//	crypted := make([]byte, len(rawData))
//	mode.CryptBlocks(crypted, rawData)
//	return crypted, nil
//}
//
////解密
//
//func AesCBCDecrypt(encryptData, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	blockSize := block.BlockSize()
//	if len(encryptData) < blockSize {
//		log.Println("encryptData is too short")
//	}
//	origData := make([]byte, len(encryptData))
//	//解加密
//	mode := cipher.NewCBCDecrypter(block, key[:blockSize])
//	mode.CryptBlocks(origData, encryptData)
//	//去除填充字符串
//	origData, err = PKCS7UnPadding(origData)
//	if err != nil {
//		return nil, err
//	}
//	return origData, nil
//}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AES加密,CBC
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

//加密base64

func Encrypt(pwd, key []byte) (string, error) {
	data, err := AesEncrypt(pwd, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

//解密base64

func Decrypt(pwd string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return "", err
	}
	dnData, errData := AesDecrypt(data, key)
	if errData != nil {
		return "", err
	}
	return string(dnData), nil
}
