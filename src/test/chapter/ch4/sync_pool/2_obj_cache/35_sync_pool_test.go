package object_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			// 这里 id 和 get出来的 不太符合预期
			// 因为 go 协程是抢占式调度
			// 循环只是负责了协程的创建，具体执行调度就不一定按顺序来了
			// 所以一般协程间通信是用 channel 的
			fmt.Printf("%d, %d\n", id, pool.Get())

			wg.Done()
		}(i)
	}

	wg.Wait()
}

//------------------------ 协程间通信，通过 channel ---------------------------- //

var wg sync.WaitGroup

func init() {
	wg = sync.WaitGroup{}
}

func TestGroutineChannel(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	ch := make(chan int, 10)

	producer(ch, &wg, pool)
	receiver(ch, &wg)

	wg.Wait()
}

func receiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Printf("-----------------------receive: %d\n", data)
				wg.Done()
			} else {
				//close(ch)
				break
			}
		}
	}()
}

func producer(ch chan int, wg *sync.WaitGroup, pool *sync.Pool) {
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			get := pool.Get().(int)
			ch <- get
			fmt.Printf("get %d, %d\n", id, get)
		}(i)
	}
}

//------------------------ 协程间通信2，通过 channel，，下面是段错误的代码，，原因还不知道😭 ---------------------------- //

func TestGroutineChannel2(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	ch := make(chan int, 10)

	producer2(ch, &wg, pool)
	wg.Add(1)
	receiver2(ch, &wg)

	wg.Wait()
}

func receiver2(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Printf("-----------------------receive: %d\n", data)
			} else {
				close(ch)
				break
			}
		}
		wg.Done()
	}()
}

func producer2(ch chan int, wg *sync.WaitGroup, pool *sync.Pool) {
	for i := 0; i < 10; i++ {
		go func(id int) {
			wg.Add(1)

			get := pool.Get().(int)
			ch <- get
			fmt.Printf("get %d, %d\n", id, get)

			wg.Done()
		}(i)
	}
}

//------------------------ 工作线程 ---------------------------- //
func workerPool(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("start job", j)
		time.Sleep(time.Second)
		fmt.Println("finish job", j)
		results <- j
	}
}

func TestWorkJob(t *testing.T) {
	const numJobs = 5
	jobs := make(chan int)    //
	results := make(chan int) //

	go workerPool(jobs, results)

	go func() {
		for r := range results {
			fmt.Println("--------------------receive :", r)
			wg.Done() //接收到数据,表示完成了一份工作
		}
	}()

	for i := 1; i <= numJobs; i++ {
		wg.Add(1) //标记开始一份工作
		jobs <- i
	}

	wg.Wait()

}
