# demo-pb-dynamic-go

一个小 demo，运行时读取 pb 结构，用运行时获取的 pb 结构打印消息

"读取 pb 结构"的原理是 protoc 可以产生 fileDescriptorSet 写入文件，起到 parser 的作用。

## dev dependency

1. protoc

不需要 protoc-go 因为我们并不会将 pb 编译成 go 文件。我们只是在运行时读取 pb 的结构。

