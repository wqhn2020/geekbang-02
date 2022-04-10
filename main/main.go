package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

// 课后练习 1.2
// 基于 Channel 编写一个简单的单协程生产者消费者模型。
// 要求如下：
// 1）队列：队列长度 10，队列元素类型为 int
// 2）生产者：每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
// 3）消费者：每2秒从队列中获取一个元素并打印，队列为空时消费者阻塞
// 4）主协程30秒后要求所有子协程退出。
// 5）要求优雅退出，即消费者协程退出前，要先消费完所有的int。

var (
	wg sync.WaitGroup // 定义一个WainGroup，用于实现主协程等待子协程结束后才退出。

	p Producer  // 生产者
	c Consumer  // 消费者
)

type Producer struct {
	Times    int // 生产的次数
	Interval int // 生产的频率，以秒为单位
}

// Consumer 消费者
type Consumer struct {
	Producer // 结构体嵌套，类似JAVA中的继承
}

// Produce 生产方法
// 每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
func (p Producer) Produce(queue chan<- int, ctx context.Context) {
	go func() {
	LOOP:
		for {
			p.Times = p.Times + 1

			val := p.Times
			queue <- val // 把生产者的生产次数作为“值”，放入通道。如果通道满了，则这里会阻塞。

			fmt.Printf("生产者: 第%d次生产, 值为：%d\n", p.Times, val)
			time.Sleep(time.Duration(p.Interval) * time.Second)

			select {
			case <-ctx.Done(): // 检查一下有没有收到main的退出通知(这里不会阻塞)
				close(queue)
				break LOOP
			default:
			}
		}
		wg.Done() // 通知main协程，子协程已退出
	}()
}

// Consume 消费方法
// 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
func (c Consumer) Consume(queue <-chan int, ctx context.Context) {
	go func() {
	LOOP:
		for {
			c.Times = c.Times + 1

			val := <-queue // 从通道中读取数据，如果通道中没有数据，则这里会阻塞
			fmt.Printf("-->消费者: 第%d次消费, 值为：%d\n", c.Times, val)
			time.Sleep(time.Duration(c.Interval) * time.Second)

			select {
			case <-ctx.Done(): // 检查一下有没有收到main的退出通知(这里不会阻塞)

				var items []int
				for val = range queue {
					items = append(items, val)
				}
				fmt.Printf("-->消费者: 最后一次消费, 值为：%v\n", items)
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
	go p.Produce(queue, ctx)

	// 开启消费者协程
	go c.Consume(queue, ctx)

	time.Sleep(time.Second * 30)
	cancel() // 30秒后，main协程通知子协程结束
	fmt.Println("main waiting")
	wg.Wait() // 等待子协程退出，这种方式比较优雅。
	fmt.Println("-- done --")
}

/*
启动命令：
$ go run main/main.go -m wb
$ go run main/main.go -m je
 */
func init() {
	// 解析程序入参，运行模式
	mode := flag.String("m", "wb", "请输入运行模式：\nwb（温饱模式）生产速度快过消费速度、\nje（饥饿模式）生产速度慢于消费速度)")
	flag.Parse()

	p = Producer{}
	c = Consumer{}

	if *mode == "wb" {
		fmt.Println("运行模式：wb（温饱模式）生产速度快过消费速度")
		p.Interval = 1     // 每隔1秒生产一次
		c.Interval = 5     // 每隔5秒消费一次
	} else {
		fmt.Println("运行模式：je（饥饿模式）生产速度慢于消费速度")
		p.Interval = 5     // 每隔5秒生产一次
		c.Interval = 1     // 每隔1秒消费一次
	}

}
