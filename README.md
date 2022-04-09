# go-chan-dealock-02

估计是通道与wg.wait()的问题，这个程序会发现死锁：

输出日志如下：





GOROOT=D:\go\go1.17.8 #gosetup
GOPATH=D:\go\go-path #gosetup
D:\go\go1.17.8\bin\go.exe build -o C:\Users\tanji\AppData\Local\Temp\GoLand\___go_build_main_go.exe E:\learning-go\go-project\geekbang-02\main\main.go #gosetup
C:\Users\tanji\AppData\Local\Temp\GoLand\___go_build_main_go.exe #gosetup
生产者: 第1次生产, 值为：1
-->消费者: 第1次消费, 值为：1
生产者: 第2次生产, 值为：2
-->消费者: 第2次消费, 值为：2
生产者: 第3次生产, 值为：3
生产者: 第4次生产, 值为：4
生产者: 第5次生产, 值为：5
-->消费者: 第3次消费, 值为：3
生产者: 第6次生产, 值为：6
-->消费者: 第4次消费, 值为：4
生产者: 第7次生产, 值为：7
生产者: 第8次生产, 值为：8
-->消费者: 第5次消费, 值为：5
生产者: 第9次生产, 值为：9
生产者: 第10次生产, 值为：10
-->消费者: 第6次消费, 值为：6
生产者: 第11次生产, 值为：11
生产者: 第12次生产, 值为：12
-->消费者: 第7次消费, 值为：7
生产者: 第13次生产, 值为：13
生产者: 第14次生产, 值为：14
-->消费者: 第8次消费, 值为：8
生产者: 第15次生产, 值为：15
生产者: 第16次生产, 值为：16
-->消费者: 第9次消费, 值为：9
生产者: 第17次生产, 值为：17
生产者: 第18次生产, 值为：18
-->消费者: 第10次消费, 值为：10
生产者: 第19次生产, 值为：19
生产者: 第20次生产, 值为：20
-->消费者: 第11次消费, 值为：11
生产者: 第21次生产, 值为：21
-->消费者: 第12次消费, 值为：12
生产者: 第22次生产, 值为：22
-->消费者: 第13次消费, 值为：13
生产者: 第23次生产, 值为：23
-->消费者: 第14次消费, 值为：14
生产者: 第24次生产, 值为：24
-->消费者: 第15次消费, 值为：15
生产者: 第25次生产, 值为：25
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x1)
	D:/go/go1.17.8/src/runtime/sema.go:56 +0x25
sync.(*WaitGroup).Wait(0x6fc23ac00)
	D:/go/go1.17.8/src/sync/waitgroup.go:130 +0x71
main.main()
	E:/learning-go/go-project/geekbang-02/main/main.go:105 +0x18a

goroutine 9 [chan send]:
main.Producer.Produce.func1()
	E:/learning-go/go-project/geekbang-02/main/main.go:37 +0x73
created by main.Producer.Produce
	E:/learning-go/go-project/geekbang-02/main/main.go:31 +0xda

Process finished with the exit code 2