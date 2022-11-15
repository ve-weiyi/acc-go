// 1.包名：包名一般与目录相同,相同目录只能有一个包名
package main

//2.导包：可以使用 import "p1" 和 import ("p1","p2" )
import (
	"errors"
	"fmt"
	"time"
)

// 3.init先于main运行
func init() {
	fmt.Println("init invoke")
}

// 4.只有main包下的main才能运行
func main() {
	fmt.Println("main invoke")
	Test7()
}

func Test7() {
	// 传参的时候显式类型像隐式类型转换，双向管道向单向管道转换
	ch := make(chan string) //无缓冲channel
	go producer(ch)         // 子go程作为生产者
	consumer(ch)            // 主go程作为消费者
}

func producer(out chan<- string) {
	for i := 0; i < 10; i++ {
		data := time.Now().String()
		time.Sleep(1 * time.Second)
		fmt.Println("生产者生产数据:", data)
		out <- data // 缓冲区写入数据
	}
	close(out) //写完关闭管道
}

func consumer(in <-chan string) {
	// 无需同步机制，先做后做
	// 没有数据就阻塞等
	for data := range in {
		fmt.Println("消费者得到数据：", data)
	}

}

func Test6() {
	//编译时异常
	errors.New("err")
	//运行时异常
	panic("err")
}

// 5.defer语句在函数运行结束后才执行
func DeferTest() {
	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")
}
