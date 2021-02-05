## Go的高并发
### Goroutine与线程的区别
* 内存消耗更少
* 创建与销毁的开销更小
  1. 线程的创建需要行操作系统申请资源，并在销毁时归还，开销较大；
  2. 协程的创建和销毁由Go语言自身管理，不需要与操作系统打交道。
* 切换的开销更小
  1. 线程的调度方式是抢占式的，如果一个线程的执行时间超过了分配给它的时间片，就会被其他可执行的线程抢占。在线程切换的过程中需要保存/恢复所有的寄存器信息，比如16个通用寄存器，PC（Program Counter）、SP（Stack Pointer）段寄存器等等。
  2. goroutine的调度是协同式的，它不会直接地与操作系统内核打交道。当goroutine进行切换的时候，之后很少量的寄存器需要保存和恢复（PC和SP）。因此goroutine的切换效率更高。
### Goroutine的调度模型
*  Go的调度没有时间片的概念，只有如下几种情况才会发生协程的切换
  1. Channel接收或者发送会造成阻塞的消息
  2. 当一个新的goroutine被创建时
  3. 可以造成阻塞的系统调用，如文件和网络操作
  4. 垃圾回收
* 三种调度器（Process、GoThread、Goroutine）
  1. 在一个Go程序中，可用的线程数是通过GOMAXPROCS来设置的，默认值是可用的CPU核数。可以用runtime包来动态改变这个值。OSThread调度在processor上，goroutines调度在OSThreads上。
  2. Golang的调度器可以利用多processor资源，在任意时刻，M个goroutine需要被调度到N个OS threads上，同时这些threads运行在至多GOMAXPROCS个processor上（N <= GOMAXPROCS）。Go scheduler将可运行的goroutines分配到多个运行在一个或多个processor上的OS threads上。
  3. 每个processor有一个本地goroutine队列。同时有一个全局的goroutine队列。每个OSThread都会被分配给一个processor。最多只能有GOMAXPROCS个processor，每个processor同时只能执行一个OSThread。Scheculer可以根据需要创建OSThread。
  4. 在每一轮调度中，scheduler找到一个可以运行的goroutine并执行直到其被阻塞。由此可见，操作系统的一个线程下可以并发执行上千个goroutine，每个goroutine所占用的资源和切换开销都很小，因此，goroutine是golang适合高并发场景的重要原因。
## channel与context
* 
## Go的垃圾回收机制
* 
## interface与interface{}
* 
## Go的反射
* 111
## make & new
* 相同：
* 不同：
1. make是只针对map、slice与channel的初始化，返回值的本身
2. new可以针对任意类型，但是返回指针，置为0值
## go slice
* append 是拓展到原来的2倍或1.25倍
1. 如果原Slice容量小于1024，则新Slice容量将扩大为原来的2倍；
2. 如果原Slice容量大于等于1024，则新Slice容量将扩大为原来的1.25倍；
## defer
1. 延迟函数的单数在defer语句出现之前就已经确定
2. 延迟函数按照后进先出的原则执行，通过栈来维护
3. 延迟函数可能操作住函数具名的返回值