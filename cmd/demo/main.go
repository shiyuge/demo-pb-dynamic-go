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
		log.Println("found file ", file.Name, " in file descriptor set")
		messages := file.GetMessageType()
		for _, m := range messages {
			log.Println("found message ", m.Name, " in file ", file.Name)
			tryToParseIncomingMessageWithDescriptor(incomingMessage, m)
			fields := m.GetField()
			for _, field := range fields {
				log.Println("found field ", field.Name, " in message ", m.Name, " in file ", file.Name)
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
		log.Println("cannot parse incoming message with message descriptor ", d.Name, " ", err.Error())
		return
	}

	log.Println("message parsed successfully: ", m.String())
}
