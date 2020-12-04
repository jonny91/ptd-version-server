package controlls

import (
	"PTDVersionServer/dto"
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

var (
	nsqProducerConfig *nsq.Config
	nsqConsumerConfig *nsq.Config
	NSQProducer       *nsq.Producer
	isInit            bool
	err               error
)

type MessageHandler struct {
}

func (h *MessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	//write to db
	msg := &dto.MissionResult{}
	err := json.Unmarshal(m.Body, msg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	DB.AutoMigrate(msg)
	DB.Create(msg)
	return nil
}

func InitMQ() error {
	nsqConfig := nsq.NewConfig()
	NSQProducer, err = nsq.NewProducer("39.105.34.250:4150", nsqConfig)
	if err != nil {
		return err
	}

	isInit = true
	return nil
}

func Send2MQ(topic string, msg []byte) {
	if isInit {
		err = NSQProducer.Publish(topic, msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ReceiveFromMQ(topic string) {
	nsqConsumerConfig := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, "channel", nsqConsumerConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	consumer.AddHandler(&MessageHandler{})

	err = consumer.ConnectToNSQLookupd("39.105.34.250:4161")
	if err != nil {
		log.Fatal(err)
	}

	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//<-sigChan
	//consumer.Stop()
}
