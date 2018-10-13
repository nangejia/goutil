package main

import (
	"github.com/nangejia/goutil/crypto/des"
	"fmt"
)

//测试des对称加密
func main() {
	cbcKey := []byte("nangejia") //秘钥长度8字节

	//des的CBC分组加密
	plainText := []byte("南哥家，使用go实现des算法的cbc分组模式的加密")           //原文
	cbcCipherText := des.DesEncryptByCBC(plainText, cbcKey)     //加密
	cbcResultText := des.DesDecryptByCBC(cbcCipherText, cbcKey) //还原密文
	fmt.Println(string(cbcResultText))                          //打原结果

	//des的CTR分组加密
	ctrKey := []byte("nangejia")                              ////秘钥长度8字节
	crtCipherText := des.DesCryptByCTR(plainText, ctrKey)     //加密
	crtResultText := des.DesCryptByCTR(crtCipherText, ctrKey) //加密
	fmt.Println(string(crtResultText))                        //打原结果
}
