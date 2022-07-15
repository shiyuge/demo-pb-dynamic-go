export LOCAL_MODULE := github.com/shiyuge/demo-pb-dynamic-go

.PHONY: fmt
fmt:
	@go fmt ./...
	@goimports -local $(LOCAL_MODULE) -w $$(find . -type f -name '*.go' -not -path "./*_gen/*" -not -path "*/mock/*")
	@go mod tidy