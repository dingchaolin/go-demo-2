package main

import (
	"crypto/rc4"
	"log"
	//"io"
	//"crypto/cipher"
	"crypto/md5"
)


func main(){
	key := "123456"
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(md5sum[:]))
	if err != nil{
		log.Fatal(err)
	}

	buf := []byte("hello")

	cipher.XORKeyStream(buf, buf)
	log.Print( string(buf) )

	//解密
	//新的作用域 不会重名
	{
		cipher, err := rc4.NewCipher([]byte(md5sum[:]))
		if err != nil{
			log.Fatal(err)
		}
		cipher.XORKeyStream(buf, buf)
		log.Printf(string(buf))
	}
}