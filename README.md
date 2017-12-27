# go-demo-2

# 1. 入门
- go语言是google2009年推出的语言
- 基于c语言 
- CSP并发模型 信道
- erlang  golang 都能充分利用多核

## 1.1 安装
- go env 查看环境信息
- go run hello.go
- go build hello.go 得到一个可执行的二进制文件
- GOOS=linux go build hello.go  elf
- GOOS=windows go build -o hello.exe hello.go exe
- GOOS=darwin go build hello.go 
- -o 指定生成的文件名
- 安卓的底层就会说linux 只是在上面跑了一个java虚拟机
- windows  set GOOS=windows
- 然后编译 就可以在windows上编译各种版本的程序了
- gofmt -w hello.go   命令 设置 代码缩进
- goimports -w hello.go 没有引入的包会自动引入
- go 中的接口跟指针类似 是指针的语法糖






