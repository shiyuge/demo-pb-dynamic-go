package main

import (
	"io"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/types/dynamicpb"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: first argument should be a compiled protobuf file descriptor set file")
	}

	descriptorFileName := os.Args[1]
	incomingMessageFileName := os.Args[2]

	run(descriptorFileName, incomingMessageFileName)
}

func run(descriptorFileName string, incomingMessageFileName string) {
	descriptorFileSet, err := readFile(descriptorFileName)
	if err != nil {
		log.Fatalf("fail to read descriptorFileSet: %+v", err)
	}

	incomingMessage, err := readFile(incomingMessageFileName)
	if err != nil {
		log.Fatalf("fail to read incomingMessage: %+v", err)
	}

	var f descriptor.FileDescriptorSet
	err = proto.Unmarshal(descriptorFileSet, &f)
	if err != nil {
		log.Fatalf("fail unmarshall FileDescriptorSet: %+v", err)
	}

	files := f.GetFile()
	for _, file := range files {
		log.Printf("found file %s in file descriptor set\n", file.GetName())
		messages := file.GetMessageType()
		for _, m := range messages {
			log.Printf("found message %s in file %s\n", m.GetName(), file.GetName())
			tryToParseIncomingMessageWithDescriptor(incomingMessage, m)
			fields := m.GetField()
			for _, field := range fields {
				log.Printf("found field %s in message %s in file %s\n", field.GetName(), m.GetName(), file.GetName())
			}
		}
	}
}

func readFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func tryToParseIncomingMessageWithDescriptor(v []byte, d *descriptor.DescriptorProto) {
	m := dynamicpb.NewMessage(d.ProtoReflect().Descriptor())
	err := proto.Unmarshal(v, m)
	if err != nil {
		log.Printf("cannot parse incoming message with message descriptor %s: %+v\n", d.GetName(), err.Error())
		return
	}

	log.Printf("message parsed successfully: %s\n", m.String())
}
