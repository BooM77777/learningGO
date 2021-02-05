package core

import (
	"errors"
	"sync"
)

// MQClient 消息队列客户端
type MQClient struct {
	topics map[string]<-chan interface{} // 订阅的只读通道
	server *MQServer

	rwMutex       sync.RWMutex
	mutexPreTopic map[string]sync.RWMutex
}

// NewMQClient 构造函数
func NewMQClient(server *MQServer) *MQClient {
	return &MQClient{
		topics:        make(map[string]<-chan interface{}),
		server:        server,
		rwMutex:       sync.RWMutex{},
		mutexPreTopic: make(map[string]sync.RWMutex),
	}
}

// GetMsg 获取队列中的消息
func (client *MQClient) GetMsg(topic string) ([]interface{}, error) {

	client.rwMutex.RLock()
	channel, ok := client.topics[topic]
	mutex, _ := client.mutexPreTopic[topic]
	client.rwMutex.RUnlock()

	if ok == false {
		return nil, errors.New("不存在这个主题")
	}

	ret := make([]interface{}, 0, 512)
	mutex.Lock()
	for len(channel) > 0 {
		ret = append(ret, <-channel)
	}
	mutex.Unlock()

	return ret, nil
}

// Publish 向一个主题推送消息
func (client *MQClient) Publish(topic string, message interface{}) error {
	err := (*client.server).publish(topic, message)
	return err
}

// Subscribe 订阅一个主题
func (client *MQClient) Subscribe(topic string) error {

	client.rwMutex.RLock()
	_, ok := client.topics[topic]
	client.rwMutex.RUnlock()

	if ok {
		return errors.New("该主题已订阅")
	}

	topicChan, err := (*client.server).subscribe(topic)
	if err != nil {
		return err
	}

	client.rwMutex.Lock()
	client.topics[topic] = topicChan
	client.mutexPreTopic[topic] = sync.RWMutex{}
	client.rwMutex.Unlock()

	return nil
}

// Unsubscribe ...
func (client *MQClient) Unsubscribe(topic string) error {

	client.rwMutex.RLock()
	channel, ok := client.topics[topic]
	client.rwMutex.RUnlock()

	if ok == false {
		return errors.New("并未订阅该主题")
	}

	err := (*client.server).unsubscribe(topic, channel)
	if err != nil {
		return err
	}

	client.rwMutex.Lock()
	delete(client.topics, topic)
	client.rwMutex.Unlock()

	return nil
}
