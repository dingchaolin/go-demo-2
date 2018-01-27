package main

import (
	"io"
	//"net"
	//"fmt"
	//"crypto/cipher"
	"crypto/rc4"
	//"crypto"
	"crypto/md5"
	"os"
	"flag"
)


type CryptoWriter struct{
	w io.Writer
	cipher *rc4.Cipher
}

func NewCryptoWriter(w io.Writer, key string) io.Writer{
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil{
		panic(err)
	}
	return &CryptoWriter{
		w: w,
		cipher:cipher,
	}
}
// 把b里面的数据进行加密 之后写入到w.w中
// 调用w.w.Write进行写入
func ( w *CryptoWriter) Write(b []byte)(int, error){
	buf := make([]byte, len(b))
	w.cipher.XORKeyStream(buf,b)
	return w.w.Write(buf)
}

type CryptoReader struct{
	r io.Reader
	cipher *rc4.Cipher
}

func NewCryptoReader(r io.Reader, key string) io.Reader{
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil{
		panic(err)
	}
	return &CryptoReader{
		r:r,
		cipher:cipher,
	}
}


func ( r *CryptoReader) Read(b []byte)(int, error){
	n, err := r.r.Read(b)
	buf := b[:n]
	r.cipher.XORKeyStream(buf,buf)
	return n, err

}
func main(){

	//key := "123456"
	//remote, err := net.Dial("tcp", "")
	//if err != nil{
	//	fmt.Println( err )
	//	return
	//}
	//w := NewCryptoWriter( remote, key)
	//w.Write([]byte("hello"))
	//
	//r := NewCryptoReader(remote, key)
	//buf := make([]byte, 1024)
	//r.Read(buf)
	// echo hello | ./crypto
	// echo hello | ./crypto ./crypto
	r := NewCryptoReader(os.Stdout, "123456")
	io.Copy(os.Stdout, r)

	w := NewCryptoWriter(os.Stdout, "123456")
	io.Copy(w, os.Stdin)

}
