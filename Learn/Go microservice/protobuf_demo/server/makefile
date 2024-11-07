.PHONY: gen help

PROTO_DIR=pb

gen:
	protoc --proto_path=pb --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative  book/book.proto book/price.proto author/author.proto
help:
	@echo "make gen - 生成pb及grpc代码"
