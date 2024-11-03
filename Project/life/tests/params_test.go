package tests

import (
	"life/pkg/models"
	"reflect"
	"testing"
)

func TestRemoveTopStruct(t *testing.T) {
	type args struct {
		fields map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "正常用例-单层结构体",
			args: args{
				fields: map[string]string{
					"User.Name": "名字不能为空",
					"User.Age":  "年龄必须大于0",
				},
			},
			want: map[string]string{
				"Name": "名字不能为空",
				"Age":  "年龄必须大于0",
			},
		},
		{
			name: "多个字段带相同结构体前缀",
			args: args{
				fields: map[string]string{
					"User.Name":     "名字不能为空",
					"User.Age":      "年龄必须大于0",
					"User.Email":    "邮箱格式不正确",
					"User.Password": "密码长度不够",
				},
			},
			want: map[string]string{
				"Name":     "名字不能为空",
				"Age":      "年龄必须大于0",
				"Email":    "邮箱格式不正确",
				"Password": "密码长度不够",
			},
		},
		{
			name: "不同结构体前缀",
			args: args{
				fields: map[string]string{
					"User.Name":    "用户名不能为空",
					"Order.ID":     "订单ID无效",
					"Product.Code": "产品编码错误",
				},
			},
			want: map[string]string{
				"Name": "用户名不能为空",
				"ID":   "订单ID无效",
				"Code": "产品编码错误",
			},
		},
		{
			name: "空map测试",
			args: args{
				fields: map[string]string{},
			},
			want: map[string]string{},
		},
		{
			name: "特殊字符测试",
			args: args{
				fields: map[string]string{
					"User.user-name": "用户名格式错误",
					"User.email@":    "邮箱格式错误",
				},
			},
			want: map[string]string{
				"user-name": "用户名格式错误",
				"email@":    "邮箱格式错误",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := models.RemoveTopStruct(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveTopStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
