package main

import (
	"github.com/nangejia/goutil/crypto/des"
	"fmt"
)
//测试des对称加密
func main() {
	key:=[]byte("nangejia")//秘钥
	plainText:=[]byte("南哥家，使用go实现des算法的cbc分组模式的加密")//原文
	cipherText := des.DesEncryptByCBC(plainText, key)//加密
	resultText := des.DesDecryptByCBC(cipherText, key)//还原密文
	fmt.Println(string(resultText))//打原结果
}