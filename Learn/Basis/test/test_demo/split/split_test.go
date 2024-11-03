package split

import (
	"reflect"
	"testing"
)

//测试

func TestSplit(t *testing.T) {
	//got := Split("我爱你", "爱")
	//want := []string{"我", "你"}
	//if !reflect.DeepEqual(got, want) {
	//	t.Errorf("want:%v got:%v\n", want, got)
	//}
	//got = Split("a:b:c", ":")
	//want = []string{"a", "b", "c"}
	//if !reflect.DeepEqual(got, want) {
	//	t.Errorf("want:%v got:%v\n", want, got)
	//}
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple":     test{"我爱你", "爱", []string{"我", "你"}},
		"multi sep1": test{"a:b:c", ":", []string{"a", "b", "c"}},
		"multi sep2": test{"a b c d", "b c", []string{"a ", " d"}},
		"Chinese":    test{"只因你太美", "只因", []string{"", "你太美"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%v falied,want:%v got:%v\n", name, tc.want, got)
			}
		})
	}
}
func BenchmarkSplit(b *testing.B) {
	//b.N不是固定的数
	for i := 0; i < b.N; i++ {
		Split("四川山清水秀，气候宜人，最喜欢四川了", "四川")
	}
}
