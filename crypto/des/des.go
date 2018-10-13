//des对称加密
package des

import (
	"crypto/des"
	"github.com/nangejia/goutil/errutil"
	"bytes"
	"crypto/cipher"
	)

//以分组为单位进行处理的密码算法称为分组密码
//常见的对称加密分组模式有：
//CBC：密码块链模式
//CTR：计数器模式

//增充分组密码(用于加密，当分组的密码长度不够时)
//plainText原文 //bloclSize块的大小
func paddingLastGroup(plainText []byte, bloclSize int) []byte {
	//最后一组需要填充的数据长度
	padNum := bloclSize - len(plainText)%bloclSize
	//创建填空的字符切片，长度为padNum，每个字节长度为padNum
	char := []byte{byte(padNum)} //将最后一位要填充的数转化ASCII码字符
	newPlain := bytes.Repeat(char, padNum)
	//填充剩余的内容，得到一个新的内容
	newText := append(plainText, newPlain...)
	return newText //返回填充好的内容
}

//去除原文的填充内容(用于解密时)
func unPaddingLastGroup(plainText []byte) []byte {
	//获取原文的长度
	lenght := len(plainText)
	//获取原文的最后一个数据(字符)
	lastChar := plainText[lenght-1]
	//将该字符转化为int型
	num := int(lastChar)
	//返回去掉填充的原文
	return plainText[:lenght-num]
}

//des的CBC分组模式加密
//plainText是原文 key：秘钥
func DesEncryptByCBC(plainText, key []byte) []byte {
	//创建一个使用des的密码接口
	block, err := des.NewCipher(key)
	//使用错误简单封装处理错误
	errutil.Error(err)
	//填空原文
	newText := paddingLastGroup(plainText, block.BlockSize())

	//创建CBC分组加密接口
	iv := []byte("12345678") //随机向量
	//返回一个块模型接口
	blockMode := cipher.NewCBCEncrypter(block, iv)

	//创建密文切片,长度与原文一样长
	cipherText := make([]byte, len(newText))
	//加密原文生成密文
	blockMode.CryptBlocks(cipherText, newText)
	return cipherText
}

//des的CBC分组模式加密
//cipherText密文 key:私钥
func DesDecryptByCBC(cipherText, key []byte) []byte {
	//创建一个使用des的密码接口
	block, err := des.NewCipher(key)
	errutil.Error(err) //处理错误

	//创建CBS分组解密接口
	iv := []byte("12345678") //随机向量
	blockMode := cipher.NewCBCDecrypter(block, iv)

	//创建原文切片，长度与密文大小一致
	plainText := make([]byte, len(cipherText))
	//解密
	blockMode.CryptBlocks(plainText, cipherText)

	//去掉可能的填充数据，将结果返回
	return unPaddingLastGroup(plainText)
}
