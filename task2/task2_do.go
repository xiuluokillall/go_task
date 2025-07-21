package task2

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/threading"
	"sync"
	"sync/atomic"
	"time"
)

func IntSpliceAdd(uint2 *uint) {
	if uint2 == nil {
		panic("cannot splice a nil uint")
	}

	*uint2 = *uint2 + 10
}

func GoRoutineTask() {

	threading.GoSafe(func() {
		for i := 0; i < 10; i = i + 2 {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	})

	threading.GoSafe(func() {
		for i := 1; i < 10; i = i + 2 {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	})

	time.Sleep(20 * time.Second)
}

/*
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/

type Shape interface {
	Area()
	Perimeter() uint64
}

type Rectangle struct {
	Length uint64
	Width  uint64
	Height uint64
}

func (r Rectangle) Area() {
	fmt.Printf("Area of rectangle is %d\n", r.Length*r.Height)
}

func (r Rectangle) Perimeter() uint64 {
	return r.Length * r.Height * r.Width
}

/**
使用组合的方式创建一个 Person 结构体，
包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	*Person
	EmployeeID int
}

func (e Employee) PrintInfo() string {
	return fmt.Sprintf("Name is %s\n Age is %d\n EmployeeID is %d", e.Name, e.Age, e.EmployeeID)
}

/*
*编写一个程序，使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/

func ChanelBaseFunc() {
	sc := sync.WaitGroup{}
	var pushChan chan int = make(chan int)
	sc.Add(2)
	go func() {
		defer sc.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("Send Value is %d\n", i)
			pushChan <- i
			time.Sleep(1 * time.Second)
		}
		close(pushChan)
	}()

	go func(ch <-chan int) {
		defer sc.Done()
		for v := range ch {
			fmt.Printf("Receive Value is %d\n", v)
			time.Sleep(2 * time.Second)
		}
	}(pushChan)

	sc.Wait()
}

/**
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
*/

func ChannelBufferFunc() {
	var c chan int = make(chan int, 100)
	sc := sync.WaitGroup{}

	sc.Add(1)
	go func(c chan<- int) {
		defer sc.Done()
		for i := 0; i < 100; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}
		close(c)
	}(c)

	timeOut := time.After(2 * time.Second)
	sc.Add(1)
	go func() {
		defer sc.Done()
		for {
			select {
			case v, ok := <-c:
				if !ok {
					fmt.Println("Channel c is closed")
					return
				}
				fmt.Printf("Receive Value is %d\n", v)
				timeOut = time.After(2 * time.Second)
			case <-timeOut:
				fmt.Println("Timeout")
				return
			default:
				fmt.Println("waiting .....")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	sc.Wait()
}

/**
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

var countNum int = 0

func counter(m *sync.Mutex) {

	m.Lock()
	defer m.Unlock()
	countNum++
}

func SyncMutexFunc() {
	sc := sync.WaitGroup{}
	//此处需要主要 必须指针传递 如果是值传递 每个goroutine之间会复制不同的锁副本 无法真正实现互斥
	m := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		sc.Add(1)
		go func() {
			defer sc.Done()
			for j := 0; j < 100; j++ {
				counter(m)
			}
		}()
	}

	sc.Wait()
	fmt.Println(countNum)
}

/*
*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
var atomicCountNum int64 = 0

func atomicCounter() {
	atomic.AddInt64(&atomicCountNum, 1)
}
func SyncMutexAtomicFunc() {
	sc := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		sc.Add(1)
		go func() {
			defer sc.Done()
			for j := 0; j < 1000; j++ {
				atomicCounter()
			}
		}()
	}

	sc.Wait()
	fmt.Println(atomicCountNum)
}
