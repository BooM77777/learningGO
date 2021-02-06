package core

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func TestClient(t *testing.T) {

	topicList := []string{"topic1", "topic2", "topic3"}

	server := NewMQServer()
	defer server.close()

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			client := NewMQClient(server)
			var err error
			err = client.Subscribe("topic1")
			if err != nil {
				log.Println(err)
			}

			err = client.Subscribe("topic2")
			if err != nil {
				log.Println(err)
			}

			err = client.Subscribe("topic3")
			if err != nil {
				log.Println(err)
			}

			time.Sleep(5 * time.Second)
			for _, topic := range topicList {
				for i := 0; i < 10; i++ {
					var sent [8]byte
					rand.Read(sent[:])
					client.Publish(topic, sent)
				}
			}
			time.Sleep(5 * time.Second)
			for _, topic := range topicList {
				msg, _ := client.GetMsg(topic)
				fmt.Println(getGID(), topic, msg)
			}

			for _, topic := range topicList {
				msg, _ := client.GetMsg(topic)
				fmt.Println("again : ", getGID(), topic, msg)
			}

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ok!")
}
