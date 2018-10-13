# goutil 日常开发时写的一些工具类

- # 密码与安全

  在开发中，我经常使用的两种主力编程语言是Java与Go，不论那种语言，都提供了许多常用的公开加密算法的工具类，这里是用Go实现的一些我用过的加密与解密工具类。

- 常用的对称加密算法

  - DES(已完成)

    - CBC：密码块链模式

      ```go
      //des的CBC分组模式加密
      //plainText是原文 key：秘钥
      func DesEncryptByCBC(plainText, key []byte) (cipherText []byte) 
      
      //des的CBC分组模式解密
      //cipherText密文 key:私钥
      func DesDecryptByCBC(cipherText, key []byte) (plainText []byte) 
      ```

    - CTR：计数器模式

      ```go
      //des的CTR分组模式加密或者解密(使用异或运算，因此只需要一个方法)
      //srcText输入源 key：秘钥
      func DesCryptByCTR(srcText, key []byte) []byte 
      ```

  - 3DES(待完成)

  - AES(待完成)

- 常用的非对称加密(待完成)

  - RSA
  - ECC(椭圆曲线)



