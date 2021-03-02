package test

import (
	"context"
	"sync"
	"testing"
	"time"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2021/3/1 下午3:03
 * @company    ：eeo.cn
 */

//并发模型 -- 主协程先退出导致自协程未执行完成
func TestCurrent(t *testing.T) {
	go func() {
		t.Log(`test`)
		time.Sleep(time.Second)
		t.Log(`goroutine`)
	}()
}

//等待组
func TestWait(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		t.Log(`wait group`)
		wg.Done()
	}()
	wg.Wait()
}

//锁，并发map
func TestCurrentMap(t *testing.T) {
	var lk sync.RWMutex
	a := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			lk.Lock()
			defer lk.Unlock()
			a[i] = i
		}(i)
	}
	t.Log(a)
	//sync.Map 请自行学习
}

//chan:协程之间的数据通信
func TestChan1(t *testing.T) {
	ch := make(chan int, 0)
	go func() {
		ch <- 1
	}()
	data := <-ch
	t.Log(data)
}

//chan:阻塞/非阻塞
func TestChan2(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 100
	go func() {
		t.Log(<-ch)
	}()
	time.Sleep(time.Second)
}

//io多路复用，监听多个通道
func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		ch1 <- 100
	}()
	go func() {
		ch2 <- `yes`
	}()
	//阻塞，等待数据
	for {
		select {
		case num := <-ch1:
			t.Log(num)
		case str := <-ch2:
			t.Log(str)
		}
	}
}

//context 相关
/*
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
*/
func TestCtx(t *testing.T) {
	ctx := context.Background()
	//withCancel
	ctx1, cancel := context.WithCancel(ctx)
	go func(c context.Context) {
		time.Sleep(time.Second * 10)
		t.Log(`这里不显示 1`)
	}(ctx1)
	cancel()
	//withTimeOut
	ctx2, out := context.WithTimeout(ctx, time.Second*2)
	go func(c context.Context) {
		time.Sleep(time.Second * 1)
		t.Log(`这里不显示 2`)
	}(ctx2)
	out()
	//withDeadline与withTimeOut类似，感兴趣的可以比较一下
	//withValue
	ctx3 := context.WithValue(ctx, `key`, `yangsen`)
	go func(c context.Context) {
		t.Log(c.Value(`key`))
	}(ctx3)
	time.Sleep(time.Second * 3)
}

//作业：交叉打印
func TestJiaocha(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8, 10}
	t.Log(a, b)
	//请用2个通道打印a,b数组，使输出结果为 1 2 3 4 5 6 ... 10
}
