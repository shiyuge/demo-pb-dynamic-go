# demo-pb-dynamic-go

一个小 demo，运行时读取 pb 结构，用运行时获取的 pb 结构打印消息。

有一个 protobuf 文件
```protobuf
syntax = "proto3";
package m;

message User {
    int32 age = 1;
    string name = 2;
}
```

和一个 `User` 序列化后对文件。这个文件的内容就是一个 `User`，其 `age` 为 12，`name` 为 "test"。
```bash
echo 'age:12\nname:"test"' | protoc --encode=m.User testdata/message.proto > output/message.pb
```

demo 程序并不是预先将 protobuf 文件编译为 go 再编译到二进制里，而是在运行时获取 pb 结构，再利用 pb 结构
解析出 pb 消息。换言之，`User`这个结构体是在运行时作为输入进到 demo 程序里来的。

protobuf 的编译器是用 c++ 实现的，protobuf go 官方库中并无 parse protobuf 的逻辑。所以需要借用 protoc 的 parser 能力。
protoc 可以解析 protobuf 文件后输出 FileDescriptorSet （一个在 descriptor.proto 中声明好的结构体）。
demo 程序将这个 FileDescriptorSet 作为运行时输入。

## 运行 demo

执行 `make test`

```
2022/07/15 22:32:10 found file m in file descriptor set
2022/07/15 22:32:10 found message User in file m
2022/07/15 22:32:10 message parsed successfully: age:12 name:"test"
2022/07/15 22:32:10 found field age in message User in file m
2022/07/15 22:32:10 found field name in message User in file m
```

## dev dependency

protoc

不需要 protoc-go 因为我们并不会将 pb 编译成 go 文件。我们只是在运行时读取 pb 的结构。



