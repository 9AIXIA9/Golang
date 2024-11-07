package main

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"oneof_demo/api"
)

// 使用field mask实现部分更新示例
func fieldMaskDemo() {
	//client
	paths := []string{"price", "info.b"} //修改字段的路径信息
	req := api.UpdateBookRequest{
		Op: "xia",
		//虽然修改了Op但是paths修改字段路径并未包含Op，所以并不会造成改变
		//要想改变那么就要在paths添加Op
		Book: &api.Book{
			Price: &wrapperspb.Int64Value{Value: 10},
			Info:  &api.Book_Info{B: 1111},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
	}

	//server
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(req.UpdateMask, generator.CamelCase)
	bookDst := make(map[string]interface{})
	//将数据读到map[string]interface{}
	//field mask-utils支持读取到结构体等
	err := fieldmask_utils.StructToMap(mask, req.Book, bookDst)
	if err != nil {
		log.Print("field mask read to struct failed", err)
	}
	log.Printf("bookDst:%v", bookDst)
}

type Book struct {
	Price int64 //如何区分默认值和 0
	//price  sql.NullInt64 // 自定义结构体 加一个字段判断是否赋值
	//price2 *int64        //通过判断指针是否为空判断
}

// 区别默认值和 0
func distinguishDefault() {
	//var b1 Book
	//b1.Price == 0
	//var b2 = &Book{Price: 0}
	//b2.Price == 0
	//protobuf通过wraps来区分
}

// wrap示例
func wrapValueDemo() {
	book := api.Book{
		Title:  "为霞尚满天",
		Author: "霞",
		Price:  &wrapperspb.Int64Value{Value: 11},
		Memo:   &wrapperspb.StringValue{Value: "哈哈哈哈哈哈哈哈"},
	}
	if book.GetPrice() == nil {
		log.Print("没有设置price")
		//没有赋值
	}
	//赋值了就用
	log.Print(book.GetPrice())
	if book.GetMemo() == nil {
		log.Print("没有设置memo")
	}
	log.Print(book.GetMemo())
}

// oneof示例
func oneofDemo() {
	//client
	req := &api.NoticeReaderRequest{
		Msg: "我好想你",
		NoticeWay: &api.NoticeReaderRequest_Phone{
			Phone: "110",
		},
	}
	//req2 := &api.NoticeReaderRequest{
	//	Msg: "我好想你",
	//	NoticeWay: &api.NoticeReaderRequest_Email{
	//		Email: "@123qq.com",
	//	},
	//}

	//server
	switch req.NoticeWay.(type) {
	case *api.NoticeReaderRequest_Email:
		noticeWithEmail(req)
	case *api.NoticeReaderRequest_Phone:
		noticeWithPhone(req)
	}
}

func main() {
	oneofDemo()
	wrapValueDemo()
	fieldMaskDemo()
}

func noticeWithEmail(n *api.NoticeReaderRequest) {
	log.Printf("someone tell you \"%v\" by email \"%v\"", n.GetMsg(), n.GetEmail())
}

func noticeWithPhone(n *api.NoticeReaderRequest) {
	log.Printf("someone tell you \"%v\" by phone \"%v\"", n.GetMsg(), n.GetPhone())
}
