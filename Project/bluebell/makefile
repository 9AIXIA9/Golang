.PHONY: all build run goTool clean help

BINARY = "bluebell"

all: goTool build

build :
 	CGO_ENABLED = 0 GOOS = linux GOARCH = amd64 go build -o ${BINARY}

run :
	go run bluebell -confFilePath ".\settings\config.yaml"

goTool :
	go fmt ./
	go vet ./

clean :
	@if [-f ${BINARY}] ; then rm ${BINARY} ; fi

help :
		@echo "make - 格式化 Go 代码,并编译生成二进制文件"
		@echo "make build -编译 Go 代码,生成二进制文件"
		@echo "make run - 直接运行Go代码"


