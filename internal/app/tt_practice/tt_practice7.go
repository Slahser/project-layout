package tt_practice

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

type Person2 struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson2(name string, salary float64) *Person2 {
	p := &Person2{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person2) backend() {
	for f := range p.chF {
		f()
	}
}

// Set salary.
func (p *Person2) SetSalary(sal float64) {
	p.chF <- func() { p.salary = sal }
}

// Retrieve salary.
func (p *Person2) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }
	return <-fChan
}

func (p *Person2) String() string {
	return "Person2 - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

type Task struct {
	tt string
	// some state
}
type Pool struct {
	Mu    sync.Mutex
	Tasks []*Task
}

func Worker1(pool *Pool) {
	for {
		pool.Mu.Lock()
		// begin critical section:
		task := pool.Tasks[0]       // take the first task
		pool.Tasks = pool.Tasks[1:] // update the pool of tasks
		// end critical section
		pool.Mu.Unlock()
		//process(task)
		fmt.Printf(task.tt)
	}
}

func Worker2(in, out chan *Task) {
	for {
		t := <-in
		//process(t)
		fmt.Printf(t.tt)
		out <- t
	}
}

func f(left, right chan int) { left <- 1 + <-right }

var resume chan int

/**
只在你需要时进行求值，同时保留相关变量资源（内存和 cpu）：这是一项在需要时对表达式进行求值的技术
通道被命名为yield和resume，这些词经常在协程代码中使用
*/
func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

type EvalFunc func(Any) (Any, Any)

/**
- 使用锁的情景：
    - 访问共享数据结构中的缓存信息
    - 保存应用程序上下文和状态信息数据
- 使用通道的情景：
    - 与异步操作的结果进行交互
    - 分发任务
    - 传递数据所有权
*/
func Tt7() {
	runtime.GOMAXPROCS(8)

	//channel 并发访问对象
	bs := NewPerson2("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed:")
	fmt.Println(bs)

	//pending, done := make(chan *Task), make(chan *Task)
	//go sendWork(pending)       // put tasks with work on the channel
	//for i := 0; i < N; i++ {   // start N goroutines to do work
	//	go Worker(pending, done)
	//}
	//consumeWork(done)          // continue with the processed tasks

	//链式goroutine
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < 100; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0      // bang!
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 100000, ongeveer 1,5 s

	//lazy generator
	resume = integers()
	fmt.Println(generateInteger()) //=> 0
	fmt.Println(generateInteger()) //=> 1
	fmt.Println(generateInteger()) //=> 2

}
