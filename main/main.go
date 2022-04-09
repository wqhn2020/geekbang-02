package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//课后练习 1.2
//基于 Channel 编写一个简单的单协程生产者消费者模型：
//
//队列：
//队列长度 10，队列元素类型为 int
//生产者：
//每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
//消费者：
//每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞

// 定义一个WainGroup，用于
var wg sync.WaitGroup

type Producer struct {
	times    int
	interval int // 生产的频率，以秒为单位
}

// Produce 生产方法
// 每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
func (p Producer) Produce(queue chan<- int, ctx context.Context) {
	go func() {
	LOOP:
		for {
			p.times = p.times + 1

			val := p.times
			queue <- val // 把生产者的生产次数作为“值”，放入通道。如果通道满了，则这里会阻塞。

			fmt.Printf("生产者: 第%d次生产, 值为：%d\n", p.times, val)
			time.Sleep(time.Duration(p.interval) * time.Second)

			select {
			case <-ctx.Done(): // 检查一下有没有收到main的退出通知(这里不会阻塞)
				// close(queue)
				break LOOP
			default:
			}
		}
		wg.Done() // 通知main协程，子协程已退出
	}()
}

// Consumer 消费者
type Consumer struct {
	// Producer // 结构体嵌套，类似JAVA中的继承  为什么这样不行 ????
	times    int
	interval int // 生产的频率，以秒为单位
}

// Consume 消费方法
// 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
func (c Consumer) Consume(queue <-chan int, ctx context.Context) {
	go func() {
	LOOP:
		for {
			c.times = c.times + 1

			val := <-queue // 从通道中读取数据，如果通道中没有数据，则这里会阻塞
			fmt.Printf("-->消费者: 第%d次消费, 值为：%d\n", c.times, val)
			time.Sleep(time.Duration(c.interval) * time.Second)

			select {
			case <-ctx.Done(): // 检查一下有没有收到main的退出通知(这里不会阻塞)

				//var items []int
				//for val = range queue{
				//	items = append(items, val)
				//}
				//fmt.Printf("-->消费者: 最后一次消费, 值为：%v\n", items)
				break LOOP
			default:

			}
		}
		wg.Done() // 通知main协程，子协程已退出
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2) // 需要等待两个协程，即：生产者、消费者

	queue := make(chan int, 10)

	// 开启生产者协程
	p := Producer{times: 0, interval: 1}
	go p.Produce(queue, ctx)

	// 开启消费者协程
	c := Consumer{times: 0, interval: 2}
	go c.Consume(queue, ctx)

	time.Sleep(time.Second * 30)
	cancel()  // 30秒后，main协程通知子协程结束
	wg.Wait() // 等待子协程退出，这种方式比较优雅。
	fmt.Println("-- done --")
}
