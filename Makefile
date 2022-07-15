export LOCAL_MODULE := github.com/shiyuge/demo-pb-dynamic-go

.PHONY: build fmt test unit_test
fmt:
	@go fmt ./...
	@goimports -local $(LOCAL_MODULE) -w $$(find . -type f -name '*.go' -not -path "./*_gen/*" -not -path "*/mock/*")
	@go mod tidy

build:
	mkdir -p output
	go build -o output/demo ./cmd/demo

test: build
	protoc --descriptor_set_out=output/message.descriptor testdata/message.proto
	echo 'age:12\nname:"test"' | protoc --encode=m.User testdata/message.proto > output/message.pb
	output/demo output/message.descriptor output/message.pb

unit_test:
	go test -mod=mod $$(go list ./... | grep -v encryptor) -cover -coverprofile=coverage.out -coverpkg=./... -gcflags all=-l
