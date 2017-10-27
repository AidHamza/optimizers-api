package messaging

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/AidHamza/optimizers-api/config"
)

type Producer struct {
	nsq *nsq.Producer
}

func NewProducer() (*Producer, error) {
	nqsConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%d", config.App.Messaging.Host, config.App.Messaging.Port), nqsConfig)
	if err != nil {
		return &Producer{}, err
	}

	return &Producer {
		nsq: producer,
	}, nil
}

func (producer *Producer) PublishMessage(topic string, message string) error {
	err := producer.nsq.Publish(topic, []byte(message))
	if err != nil {
		return err
	}

	return nil
}