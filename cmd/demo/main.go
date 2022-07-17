package main

import (
	"io"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: demo PROTOBUF_FILE_DESCRIPTOR_SET_FILENAME INCOMING_MESSAGE_FILENAME")
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

	var f descriptorpb.FileDescriptorSet
	err = proto.Unmarshal(descriptorFileSet, &f)
	if err != nil {
		log.Fatalf("fail unmarshall FileDescriptorSet: %+v", err)
	}

	files := f.GetFile()
	for _, fileProto := range files {
		file, err := protodesc.NewFile(fileProto, nil)
		if err != nil {
			log.Fatalf("fail to create file for %s: %+v", fileProto.GetName(), err)
		}

		log.Printf("found file %s in file descriptor set\n", file.Name())
		messages := file.Messages()
		for i := 0; i < messages.Len(); i++ {
			m := messages.Get(i)
			log.Printf("found message %s in file %s\n", m.Name(), file.Name())
			tryToParseIncomingMessageWithDescriptor(incomingMessage, m)
			fields := m.Fields()
			for j := 0; j < fields.Len(); j++ {
				field := fields.Get(j)
				log.Printf("found field %s in message %s in file %s\n", field.Name(), m.Name(), file.Name())
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

func tryToParseIncomingMessageWithDescriptor(incomingMessage []byte, messageDescriptor protoreflect.MessageDescriptor) {
	m := dynamicpb.NewMessage(messageDescriptor)
	err := proto.Unmarshal(incomingMessage, m)
	if err != nil {
		log.Printf("cannot parse incoming message with message descriptor %s: %+v\n",
			messageDescriptor.Name(), err.Error())
		return
	}

	log.Printf("message parsed successfully: %s\n", m.String())
}
