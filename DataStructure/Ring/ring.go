package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type (
	//SpinLock 自旋锁
	spinLock uint32

	//Kfifo 无锁队列（暂时不考虑）
	Kfifo struct {
		buf  []int     // 缓冲区
		lock *spinLock // 自旋锁
	}
)

//Lock 请求锁
func (l *spinLock) Lock() {
	// 当请求不到时，循环请求
	// 如果是互斥锁，当请求不到时，睡眠
	for !atomic.CompareAndSwapUint32((*uint32)(l), 0, 1) {
		// CAS算法是一种有名的无锁算法。无锁编程，即不使用锁的情况下实现多线程之间的变量同步
		// 也就是在没有线程被阻塞的情况下实现变量的同步，所以也叫非阻塞同步（Non-blocking Synchronization）。
		// 需要读写的内存值V、进行比较的值A、拟写入的新值B
		// 当且仅当 V 的值等于 A时，CAS通过原子方式用新值B来更新V的值，否则不会执行任何操作。
		// 一般情况下是一个自旋操作，即不断的重试。
		runtime.Gosched() //这个函数的作用是让当前goroutine让出CPU，好让其它的goroutine获得执行的机会
	}
}

//Unlock 释放锁
func (l *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(l), 0) // 原子操作，置0
}

//NewSpinLock 构造自旋锁
func NewSpinLock() sync.Locker {
	var l spinLock
	return &l
}

//Push 向其中加入数据
func (q *Kfifo) Push(inputBuf []int) {
	q.lock.Lock()
}

func (q *Kfifo) push(inputBuf []int) {

}

func NewKfifo(len int) {

}

func main() {

}
