package utils

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// RegCodeCreat 注册码生成
func RegCodeCreat() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vCode
}

// AesEncryptByECBBase64 aes加密
func AesEncryptByECBBase64(data, key string) string {
	// 判断key长度
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic("key长度必须是 16、24、32 其中一个")
	}
	// 密钥和待加密数据转成[]byte
	originByte := []byte(data)
	keyByte := []byte(key)
	// 创建密码组，长度只能是16、24、32字节
	block, _ := aes.NewCipher(keyByte)
	// 获取密钥长度
	blockSize := block.BlockSize()
	// 补码
	originByte = pkcs5Padding(originByte, blockSize)
	// 创建保存加密变量
	encryptResult := make([]byte, len(originByte))
	// CEB是把整个明文分成若干段相同的小段，然后对每一小段进行加密
	for bs, be := 0, blockSize; bs < len(originByte); bs, be = bs+blockSize, be+blockSize {
		block.Encrypt(encryptResult[bs:be], originByte[bs:be])
	}
	return base64.StdEncoding.EncodeToString(encryptResult)
}

// pkcs5补码算法
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// ChangeChar 前两个字符和后两个字符交换
func ChangeChar(s string) string {
	b := []byte(s)
	for i := 0; i < 2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return string(b)
}

func AesDecryptByECBBase64(data, key string) string {
	// 判断key长度
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic("key长度必须是 16、24、32 其中一个")
	}
	// 反解密码base64
	originByte, _ := base64.StdEncoding.DecodeString(data)
	// 密钥和待加密数据转成[]byte
	keyByte := []byte(key)
	// 创建密码组，长度只能是16、24、32字节
	block, _ := aes.NewCipher(keyByte)
	// 获取密钥长度
	blockSize := block.BlockSize()
	// 创建保存解密变量
	decrypted := make([]byte, len(originByte))
	for bs, be := 0, blockSize; bs < len(originByte); bs, be = bs+blockSize, be+blockSize {
		block.Decrypt(decrypted[bs:be], originByte[bs:be])
	}
	// 解码
	return string(pkcs5UnPadding(decrypted))
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// PasswordEncryption 注册码加密
func PasswordEncryption(code string) string {
	//交换字符串前后位置
	swapFrRearPosition := ChangeChar(code)
	//添加公共前缀
	commonPrefix := "invited"
	AddPublicPrefix := commonPrefix + swapFrRearPosition
	//aes加密
	key := "1234567812345678"
	data := AesEncryptByECBBase64(AddPublicPrefix, key)
	return data
}

// PasswordDecryption 注册码解密
func PasswordDecryption(code string) string {
	//aes解密
	key := "1234567812345678"
	decCode := AesDecryptByECBBase64(code, key)
	//删除公共前缀
	commonPrefix := "invited"
	deletePublicPrefix := strings.TrimPrefix(decCode, commonPrefix)
	//交换字符前后位置
	swapFrRearPosition := ChangeChar(deletePublicPrefix)
	return swapFrRearPosition
}
