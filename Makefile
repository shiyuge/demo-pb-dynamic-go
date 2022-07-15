export LOCAL_MODULE := github.com/shiyuge/demo-pb-dynamic-go

.PHONY: build fmt test unit_test
fmt:
	@go fmt ./...
	@goimports -local $(LOCAL_MODULE) -w $$(find . -type f -name '*.go' -not -path "./*_gen/*" -not -path "*/mock/*")
	@go mod tidy

build:
	mkdir -p output
	go build -o output/main

test: build
	protoc --descriptor_set_out=output/message.descriptor testdata/message.proto
	output/main output/message.descriptor

unit_test:
	go test -mod=mod $$(go list ./... | grep -v encryptor) -cover -coverprofile=coverage.out -coverpkg=./... -gcflags all=-l
