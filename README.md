# 极客时间云原生训练营 - 课后练习 1.2（加强版）
基于 Channel 编写一个简单的单协程生产者消费者模型。
要求如下：
1）队列：队列长度 10，队列元素类型为 int
2）生产者：每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
3）消费者：每2秒从队列中获取一个元素并打印，队列为空时消费者阻塞
4）主协程30秒后要求所有子协程退出。
5）要求优雅退出，即消费者协程退出前，要先消费完所有的int。

## 启动命令1：go run main/main.go -m wb

```
$ go run main/main.go -m wb
运行模式：wb（温饱模式）生产速度快过消费速度
生产者: 第1次生产, 值为：1
-->消费者: 第1次消费, 值为：1
生产者: 第2次生产, 值为：2
生产者: 第3次生产, 值为：3
生产者: 第4次生产, 值为：4
生产者: 第5次生产, 值为：5
-->消费者: 第2次消费, 值为：2
生产者: 第6次生产, 值为：6
生产者: 第7次生产, 值为：7
生产者: 第8次生产, 值为：8
生产者: 第9次生产, 值为：9
生产者: 第10次生产, 值为：10
-->消费者: 第3次消费, 值为：3
生产者: 第11次生产, 值为：11
生产者: 第12次生产, 值为：12
生产者: 第13次生产, 值为：13
-->消费者: 第4次消费, 值为：4
生产者: 第14次生产, 值为：14
-->消费者: 第5次消费, 值为：5
生产者: 第15次生产, 值为：15
-->消费者: 第6次消费, 值为：6
生产者: 第16次生产, 值为：16
main waiting
生产者: 第17次生产, 值为：17
-->消费者: 最后一次消费, 值为：[7 8 9 10 11 12 13 14 15 16 17]
-- done --
```

## 启动命令2：go run main/main.go -m je
```
$ go run main/main.go -m je
运行模式：je（饥饿模式）生产速度慢于消费速度
生产者: 第1次生产, 值为：1
-->消费者: 第1次消费, 值为：1
生产者: 第2次生产, 值为：2
-->消费者: 第2次消费, 值为：2
生产者: 第3次生产, 值为：3
-->消费者: 第3次消费, 值为：3
生产者: 第4次生产, 值为：4
-->消费者: 第4次消费, 值为：4
生产者: 第5次生产, 值为：5
-->消费者: 第5次消费, 值为：5
生产者: 第6次生产, 值为：6
-->消费者: 第6次消费, 值为：6
main waiting
-->消费者: 第7次消费, 值为：0
-->消费者: 最后一次消费, 值为：[]
-- done --
```
