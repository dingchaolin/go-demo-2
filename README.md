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

## 5.3 文件的读取方式
- file.Read
- ioutil.ReadFile
- bufio.Scanner
- bufio.Reader
- io.Copy

# 6 interface
- go中可以为任何类型绑定方法，但是这个类型必须是用type声明的
- 任何内置类型都不能绑定方法

# 6.1 可见性
- 通过首字母大小写来控制可见性
- 可见性是package级别的
- 只要同一个package名 package内就可以访问该package内的所有内容，
- 包括共有属性和私有属性，不同文件也没关系，只要package名相同即可

## 6.2 必须实现接口中的所有方法才算是实现接口

## 6.3 空接口可以接受任何类型

## 6.4  tar 打包 压缩
- tar 是顺序读写
- zcap 可以读取压缩文件

## 7 协程
- 进程是管理资源的：内存 句柄 网络socket 信号 
- 线程是在一个进程中，多个线程共享进程资源
- 进程是资源管理的单位，线程是调度的单位
- 一个进程中的线程是可以占用多核的
- c10k问题
- linux下，一个线程资源至少消耗2M内存，不算调度上的开销
- io复用模型解决并发问题 redis
- 协程 本质是用户态的线程   
- 用户态 内核态  权限不同
- 调用内核中的东西，就是内核态 
- 用户态 就是用户能做的一些事
- 线程被内核挂起 就是处于内核态
- 被唤醒就是处于用户态
- 线程多了调度就比较费时
- 协程的调度是不用进内核的 调度器就是go中runtime
- 一个协程初始化4k的内存
- 调度灵活 在用户态
- 类似线程的运行方式
- 并行处理
- main函数所在的协程是主协程
- 主函数一旦退出 所有的协程都会退出

## 7.1 不带缓冲的channel 相当于打电话 
- 相当于打电话 打的时候必须已经有人在接听状态 所以必须是2个携程中才能进行
- 可以channel的大小是1 一个在用channel的时候 另外一个是不能使用的
- 一直没有人使用 会造成死锁
-发送和 接收是异步的


## 7.2 带缓冲的channel
- 相当于信箱 信箱没满的时候 有地方缓存 什么时候取都可以
- 同一协程中的channel通信 必须使用带缓冲的信道
- 两个协程进行通信，可以使用非缓冲信道

## 9. 网络编程
- 客户端和服务端
- tcp or udp
- 长连接 or 短链接
- 文本协议 or 二进制协议
- http 文本协议  https 二进制协议
- dig baidu.com 查看所有的dns ip
- 单工 电视机 收音机 都是单工的 别人给放啥 就只能看啥
- 双工 电话 
- 半双工  对讲机 同时只能一方说话
- EOF 在网络编程中 表示对方关闭了链接 唯一判断对方关闭了链接

## 9.1 服务端
- 监听端口
- 接受新的连接
- 启动协程
- 发送接手数据
- 断开连接
- 哪里阻塞go哪里

## 9.2 聊天服务

## 9.3

### tcp特性
- 面向连接 逻辑上的连接
- 可靠传输 只要连接没有断 数据一定能送达
- 有序 接受的顺序一定跟发送的顺序是一致的
- 流量控制 - tcp能控制每次发多少 每次最少发一个字节（数据需要多次取完）

## groupcache
## kcp

## 10 代理
- 监听地址
- 接受连接
- 建立到目的服务器的连接
- 数据拷贝
- 关闭连接

## 11 加密
- 非对称 公钥加密 私钥解密     或者 私钥加密 公钥解密  
- 公钥私钥都可以加密 但是不可以用相同的钥匙解密
- 对称加密  加密解密使用相同的钥匙

## 接口是方法的集合 有方法列表
- Writer 接口只有一个方法 就是 Write
- 只要实现了接口里的所有的方法 就是这个接口

## 12 单元测试- 功能测试
- 假设要测试的文件名是 crypto.go  那么测试文件的文件名是 crypto_test.go
- 函数名以 Test 开头  
```
import "testing"

func TestCryptoReader_Read(t *testing.T) {
	
}
```
- 写完Test之后会自弹出要测试的函数
- 入参必须是  t *testing.T

## 13 基准测试 Benchmark开头 - 性能测试
```
 func BenchmarkNewCryptoWriter(b *testing.B) {
	 
 }
```

