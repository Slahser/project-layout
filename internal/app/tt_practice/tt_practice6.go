package tt_practice

import (
	"fmt"
	"strconv"
	"time"
)

func t_timer() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
	close(ch)
}

func getData(ch chan string) {
	//判断channel是否阻塞
	for input := range ch {
		fmt.Printf("%s ", input)
	}

	//for {
	//	//判断channel是否阻塞
	//	input, open := <-ch
	//	if !open {
	//		break
	//	}
	//	fmt.Printf("%s ", input)
	//}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		//信息按照箭头的方向流动
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

	fmt.Println("In app()")
	go LongWait()
	go ShortWait()
	fmt.Println("About to sleep in app()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(1e9)
	fmt.Println("At the end of app()")

	//basic channel
	//buf_ch1 := make(chan string, buf)
	ch := make(chan string, 2)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
	fmt.Println()

	//有 2 个通道 ch1 和 ch2，三个协程 pump1()、pump2() 和 suck()。
	//这是一个典型的生产者消费者模式。在无限循环中，ch1 和 ch2 通过 pump1() 和 pump2() 填充整数；
	//suck() 也是在无限循环中轮询输入的，
	//通过 select 语句获取 ch1 和 ch2 的整数并输出。
	//选择哪一个 case 取决于哪一个通道收到了信息。程序在 app 执行 1 秒后结束。
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Pump1(ch1)
	defer close(ch1)
	go Pump2(ch2)
	defer close(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e8)

}
