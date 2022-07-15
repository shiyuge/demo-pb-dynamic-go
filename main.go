package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/types/dynamicpb"
)

func main() {
	name := "MyStruct"
	var packet []byte
	var pbFileProtoSet []byte
	var f descriptor.FileDescriptorSet
	e := proto.Unmarshal(pbFileProtoSet, &f)
	if e != nil {
		// handle error
	}

	files := f.GetFile()
	for _, file := range files {
		messages := file.GetMessageType()
		for _, m := range messages {
			if m.GetName() == name {
				printVal(m, packet)
			}
		}
	}

}

func printVal(desc *descriptor.DescriptorProto, v []byte) {
	m := dynamicpb.NewMessage(desc.ProtoReflect().Descriptor())
	e := proto.Unmarshal(v, m)
	if e != nil {
		// handler error
	}
}
