package main

import (
	"bytes"
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type KafkaToS3 struct {
	Consumer *Consumer
	S3Client *s3.S3
}

func NewKafkaToS3() (*KafkaToS3, error) {
	consumer, err := NewConsumer()
	if err != nil {
		return nil, err
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(s3Session)

	return &KafkaToS3{
		Consumer: consumer,
		S3Client: s3Client,
	}, nil
}

func (k *KafkaToS3) ProcessMessages(ctx context.Context, topic string, bucketName string) {
	if err := k.Consumer.Subscribe(topic); err != nil {
		log.Printf("Failed to subscribe to topic %s: %s", topic, err)
		return
	}

	for {
		select {
		case msg := <-k.Consumer.KafkaConsumer.Events():
			switch e := msg.(type) {
			case *kafka.Message:
				log.Printf("Processing message: %s", string(e.Value))

				_, err := k.S3Client.PutObject(&s3.PutObjectInput{
					Bucket: aws.String(bucketName),
					Key:    aws.String(time.Now().Format("2006-01-02-15-04-05") + ".txt"),
					Body:   bytes.NewReader(e.Value),
				})
				if err != nil {
					log.Printf("Failed to upload to S3: %s", err)
				}
			case kafka.Error:
				log.Printf("Kafka error: %v", e)
			}

		case <-ctx.Done():
			log.Println("Shutting down KafkaToS3")
			return
		}
	}
}

func (k *KafkaToS3) Close() {
	k.Consumer.Close()
}
