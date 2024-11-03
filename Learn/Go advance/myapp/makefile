
# Go项目 Makefile 脚手架

# 变量定义
BINARY_NAME=myapp
VERSION?=1.0.0
BUILD_DIR=build
MAIN_PACKAGE=./cmd/myapp

# Go相关变量
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOVET=$(GOCMD) vet
GOFMT=gofmt

# 编译标记
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# 伪目标
.PHONY: all build clean test coverage deps lint fmt vet run help

all: clean build test

build:
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

deps:
	$(GOGET) ./...
	$(GOMOD) tidy

lint:
	golangci-lint run

fmt:
	$(GOFMT) -s -w .

vet:
	$(GOVET) ./...

run:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	./$(BUILD_DIR)/$(BINARY_NAME)

help:
	@echo "make - 清理、构建、测试"
	@echo "make build - 构建应用"
	@echo "make clean - 清理构建文件"
	@echo "make test - 运行测试"
	@echo "make coverage - 生成测试覆盖率报告"
	@echo "make deps - 检查并安装依赖"
	@echo "make lint - 运行 linter"
	@echo "make fmt - 格式化代码"
	@echo "make vet - 运行 go vet"
	@echo "make run - 构建并运行应用"
