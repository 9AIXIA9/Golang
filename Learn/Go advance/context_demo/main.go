package main

//context的基础使用
import (
	"context"
	"fmt"
	"log"
	"time"
)

const (
	valueData = "xia"
	valueKey  = "name"
)

func main() {
	//创建Context
	//father := context_demo.Background() //默认值
	//context_demo.TODO()                 //也用于初始化，但是只在不确定使用哪种上下文时使用
	contextValueDemo()     //context with value 基本使用方法
	timeOutDemo()          //context处理超时问题
	cancelGoRoutinesDemo() //使用Context处理进程控制问题
}
func contextValueDemo() {
	vCtx := NewContextWithValue()
	value := vCtx.Value(valueKey)
	log.Printf("value data:%v", value)
	//不建议使用context值传递关键参数
	//关键参数应该显示的声明出来，不应该隐式处理
	//context中最好是携带签名、trace_id这类值
	//因为携带value也是key、value的形式
	//为了避免context因多个包同时使用context而带来冲突
	//key建议采用内置类型（int,bool）
	//上面的例子我们获取valueKey是直接从当前ctx获取的
	//实际我们也可以获取父context中的value
	//在获取键值对时，我们先从当前context中查找
	//没有找到会在从父context中查找该键对应的值直到在某个父context中返回 nil 或者查找到对应的值。
	//context传递的数据中key、value都是interface类型
	//这种类型编译期无法确定类型，所以不是很安全，所以在类型断言时别忘了保证程序的健壮性。
}

// NewContextWithValue 生成一个带着值的value
func NewContextWithValue() context.Context {
	ctx := context.WithValue(context.Background(), valueKey, valueData)
	//从中派生的任何context都会获取此值
	return ctx
}

func doSth(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
		}
	}
}

func timeOutDemo() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel() //这个函数是用来主动超时，一般不会使用，或者是出错误时使用
	doSth(ctx)
	//通常健壮的程序都是要设置超时时间的，避免因为服务端长时间响应消耗资源
	//一般采用withTimeout或者withDeadline来做超时控制
	//当一次请求到达我们设置的超时时间，就会及时取消，不在往下执行
	//withTimeout和withDeadline作用是一样的，就是传递的时间参数不同而已
	// 他们都会通过传入的时间来自动取消Context
	//他们都会返回一个cancelFunc方法，通过调用这个方法可以达到提前进行取消
	//不过在使用的过程还是建议在自动取消后也调用cancelFunc去停止定时减少不必要的资源浪费
}

func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 4*time.Second)
}

func cancelGoRoutinesDemo() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("I want to laugh")
	go laugh(ctx) //需要使用线程
	time.Sleep(5 * time.Second)
	fmt.Println("I want to stop")
	cancel()
	fmt.Println("I cancel laugh")
	//日常业务开发中我们往往为了完成一个复杂的需求会开多个goroutine去做一些事情
	//这就导致我们会在一次请求中开了多个goroutine确无法控制他们
	//这时我们就可以使用withCancel来衍生一个context传递到不同的goroutine中
	//当我想让这些goroutine停止运行，就可以调用cancel来进行取消
}

func laugh(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("我要闭嘴了")
			return
		default:
			fmt.Println("红红火火恍恍惚惚")
		}
	}
}
