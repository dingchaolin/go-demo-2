# 安装
- https://github.com/google/protobuf/releases
- glide get github.com/golang/protobuf/protoc-gen-go

message Person {
    int32 id = 1;
    string name = 2;
    string email = 3;
    repeated PhoneNumber phones = 4;
}

1, 2, 3, 4 是 编号 不能重复
repeated 表示列表

syntax = "proto3"  默认是 proto2

编译
./protoc-3/bin/protoc --go_out=. addressbook.proto