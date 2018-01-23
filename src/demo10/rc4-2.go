package main

import (
	"crypto/rc4"
	"log"
	"io"
	//"crypto/cipher"
	"crypto/md5"
	"os"
	"flag"
)

var (
	key = flag.String("k", "", "secret key")
)
func crypto( w io.Writer,r io.Reader, key string){
	// 创建cipher
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher([]byte(md5sum[:]))
	if err != nil{
		log.Fatal(err)
	}
	buf := make([]byte, 4096)

	// 创建buf
	for{
		// 从r里面读取数据到buf
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}



		// 加密buf
		cipher.XORKeyStream(buf[:n], buf[:n])
		// 把buf写入到w里面
		w.Write(buf[:n])
	}
}
func main(){
	flag.Parse()
	crypto(os.Stdout, os.Stdin, /*"123456"*/*key)
}
// 加密
// go run -k 123 rc4-2.go < proxy.go > proxy.go.txt
// go run rc4-2.go -k 123 < proxy.go > proxy.go.txt
// tar czf - *.go | ./rc4 -k 123 > go.tar.gz  //加密后压缩
// echo hello ./rc4 //直接加密

// 解密
// go run rc4-2.go < proxy.go.txt > proxy.go.log
// go run rc4-2.go -k 123 < proxy.go.txt > proxy.go.txt.log
// ./rc4 -k 123 < go.tar.gz | tar tzf - // 解密后解缩
// echo hello ./rc4 ./rc4 // 加密后解密