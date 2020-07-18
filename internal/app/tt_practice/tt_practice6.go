package tt_practice

import (
	"fmt"
	"strconv"
	"time"
)

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf(strconv.Itoa(v))
		case v := <-ch2:
			fmt.Printf(strconv.Itoa(v))
		}
	}
}

/**
goroutine
*/
func Tt6() {
	Trace("tt_practice6#Tt6()")
	defer Untrace("tt_practice6#Tt6()")

	fmt.Println("In main()")
	go LongWait()
	go ShortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(1e9)
	fmt.Println("At the end of main()")

	//有 2 个通道 ch1 和 ch2，三个协程 pump1()、pump2() 和 suck()。
	//这是一个典型的生产者消费者模式。在无限循环中，ch1 和 ch2 通过 pump1() 和 pump2() 填充整数；
	//suck() 也是在无限循环中轮询输入的，通过 select 语句获取 ch1 和 ch2 的整数并输出。
	//选择哪一个 case 取决于哪一个通道收到了信息。程序在 main 执行 1 秒后结束。

	ch1 := make(chan int)
	ch2 := make(chan int)

	go Pump1(ch1)
	go Pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e8)

}
