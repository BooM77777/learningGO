package core

import (
	"errors"
	"log"
	"sync"
	"time"
)

// MQServer ...
type MQServer struct {
	exit chan bool // 退出控制

	topics map[string][]chan interface{} // 订阅者通道

	cap int // 每一个新创建通道的容量

	// 同步控制
	lock         sync.RWMutex            //全局锁
	lockPerTopic map[string]sync.RWMutex // 每一个主题一个锁
}

// NewMQServer 构造函数
func NewMQServer() *MQServer {
	return &MQServer{
		exit:         make(chan bool),
		topics:       make(map[string][]chan interface{}),
		cap:          256,
		lock:         sync.RWMutex{},
		lockPerTopic: make(map[string]sync.RWMutex),
	}
}

func (c *MQServer) isRunning() bool {
	select {
	case <-c.exit:
		log.Println("消息队列已关闭")
		return false
	default:
		return true
	}
}

// publish 推送消息
func (c *MQServer) publish(topic string, message interface{}) error {
	if c.isRunning() {
		c.lock.RLock()
		subscribers, ok := c.topics[topic]
		c.lock.RUnlock()
		if ok == false {
			return errors.New("没有这个主题")
		}
		c.broadcast(topic, subscribers, message)
	}
	return nil
}

func (c *MQServer) broadcast(topic string, subscribers []chan interface{}, message interface{}) {

	// 定时
	idleDuration := 5 * time.Millisecond
	idleTimeout := time.NewTimer(idleDuration)
	defer idleTimeout.Stop()

	c.lock.RLock()
	mutex, _ := c.lockPerTopic[topic]
	c.lock.RUnlock()

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		for _, subscriber := range subscribers {
			idleTimeout.Reset(idleDuration)
			select {
			// case <-idleTimeout.C: // 超时
			case <-c.exit: //提前退出
				return
			default:
				subscriber <- message
			}
		}
	}()

	return
}

// subscribe 订阅
func (c *MQServer) subscribe(topic string) (<-chan interface{}, error) {

	if c.isRunning() {
		ch := make(chan interface{}, c.cap)

		c.lock.Lock()
		if _, ok := c.topics[topic]; !ok {
			c.topics[topic] = make([]chan interface{}, 0, 1024)
			c.lockPerTopic[topic] = sync.RWMutex{}
		}
		mutex, _ := c.lockPerTopic[topic]
		c.lock.Unlock()

		mutex.Lock()
		c.topics[topic] = append(c.topics[topic], ch)
		mutex.Unlock()

		return ch, nil
	}
	return nil, errors.New("消息队列已关闭")
}

func (c *MQServer) unsubscribe(topic string, subscriber <-chan interface{}) error {
	if c.isRunning() {

		c.lock.RLock()
		subscribers, ok := c.topics[topic]
		mutex, _ := c.lockPerTopic[topic]
		c.lock.RUnlock()

		if ok == false {
			return errors.New("订阅者不存在")
		}
		// 异步删除 subscriber
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			for i := range subscribers {
				if subscribers[i] == subscriber {
					subscribers[i] = subscribers[len(topic)-1]
					subscribers = subscribers[:len(topic)-1]
					return
				}
			}
		}()

		return nil
	}
	return errors.New("消息队列已关闭")
}

func (c *MQServer) close() {
	// 发送关闭信息
	close(c.exit)
	// 将所有还为推送的消息进行推送
}
