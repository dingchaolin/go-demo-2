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
- goTrace

# 2. 语法

## 2.1
- go通过package组织
- package 关键字
- 放在程序的第一行
- 两种package 一种是库package 一种二进制package
- 二进制package使用main来表示 库package的名字跟go文件所在目录的名字一致
- 同一个目录下的go文件只有一个package名
- 同一个目录下的main package， 只能有一个main函数
- go build 文件名  生成一个可执行文件 如果同一个目录下有多个文件 会把所有的文件一起编译
- go install  -> go build , 把生成的文件诺挪到bin目录下
- 通过关键字import引入package
- 过个package可以使用括号括起来
- 引入但是没有使用的package会报错
- go run 指针对的单个go文件
- go build 和 go install 是针对package级别的 一个package会有多个go文件
- no install location for directory
- 1. 目录必须在gopath下的src下
- 2. package全路径是src为根的路径

## 2.2
- 代码风格只有一个风格 gofmt
- := 局部变量才能使用
- *T 即为类型T的指针
- &t 即为取变量t的地址
- *p 即为取指针变量所指向的内容
- go为了安全考虑 不允许将十六进制的数字赋值为给一个指针类型
- &变量  的类型就是 *T 

## 2.3 gotty
- 一个终端分享程序 让终端在浏览器中打开
- https://github.com/yudai/gotty/releases 下载
- tar zxvf gotty_2.0.0-alpha.3_darwin_amd64.tar
- ./gotty -p 8080 -w bash 启动


## 2.4 lshell 
- https://github.com/ghantoos/lshell.git
- 启动之后会进入一个新的shell 
- 只能执行 ls ll echo 类似的命令

# 3. 语法
- 整数 int int32 int64 uint uint32 uint64 int8 int16 
- uint8  byte  等价的
- xxd 把一个文件的ascii打印出来 一边是16进制 一遍是实际字符
- xxd 文件名
- 字符串不可修改
- 通过跟[]byte相互转化来修改
- bool true false
- linux 下 proc目录下 所有的数字目录都是进程id

# 4. 切片
- 切片是对数组某个数据范围的引用
- 数组中的内存地址是连续的
- 切片切的长度的时间复杂度都是一样的
- 2步 1.记录切的起始地址 2.记录长度
- 空切片跟 nil相等
- 切片被赋值之后 相当于两个切片是等价的

## 4.1 map
- hash方式的
- 无序的
- O(1)的访问时间
- map的遍历顺序不确定
- log.Fatal(err) 打印错误 并直接退出程序
- 结构体的成员变量都是连续内存
- 结构体数组也都是连续内存
- go中堆和栈没有明确区分
- 栈是函数内的一块内存 是一个局部作用
- 堆是一个全局的内存块 不会随着函数的结束而结束
- go的一个携程的栈是4k
- map 比struct要占内存
- string 占16个字节

## 4.2 序列化 反序列化
- 序列化 把内存中的对象编程字符串的形式放到磁盘上
- 反序列化 把磁盘中的字符串取出作为对象放在内存中
- rune -> int32

# 5 函数
- echo $? 查看退出错误码
- 函数要素  
- 函数名 参数 返回值
- 如果两个函数的参数跟返回值都是一样的则任务是一个类型的
- 两个函数的相等与否与名字无关
- defer 函数返回前执行

## 5.1 error的处理方式
### 1. 出错退出进程 - 初始化的时候
- 比如初始化进程数据的时候，如果失败直接打印错误并退出

### 2. 重试的时候， 比如http请求
- 如果err != nil 就continue 接着重试 直至重试次数用完

### 3. 函数中，把err返回/上抛
- 函数中不处理err ，把err上抛， 让调用者处理

## 5.2 panic - defer recover
- panic - defer recover 让当前函数退出





