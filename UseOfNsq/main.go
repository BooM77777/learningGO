package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/nsqio/go-nsq"
)

func main() {
	p, err := NewProducer("127.0.0.1:4150")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer p.Stop()

	for i := 0; i < 10; i++ {
		p.Publish("test", "hello!!!"+strconv.Itoa(i))
	}
	fmt.Println("done!")
}

// Producer 生产者
type Producer struct {
	producer *nsq.Producer
}

// NewProducer 构造函数
func NewProducer(addr string) (*Producer, error) {
	p, err := InitProducer(addr)
	return &Producer{producer: p}, err
}

// Publish 发送消息
func (p *Producer) Publish(topic string, message string) (err error) {

	if message == "" {
		return errors.New("message is empty")
	}

	if err = p.producer.Publish(topic, []byte(message)); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Stop 关闭
func (p *Producer) Stop() {
	p.producer.Stop()
}

// InitProducer 初始化生产者
func InitProducer(addr string) (p *nsq.Producer, err error) {

	var config *nsq.Config

	config = nsq.NewConfig()
	if p, err = nsq.NewProducer(addr, config); err != nil {
		return nil, err
	}
	return p, nil
}
