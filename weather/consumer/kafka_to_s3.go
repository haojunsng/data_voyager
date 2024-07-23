package main

import (
	"bytes"
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/confluentinc/confluent-kafka-go/kafka"
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
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatalf("Failed to create S3 session: %s", err)
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
		log.Fatalf("Failed to subscribe to topic: %s", err)
		return
	}

	for {
		select {
		case msg := <-k.Consumer.KafkaConsumer.Messages():
			log.Printf("Processing message: %s", string(msg.Value))

			_, err := k.S3Client.PutObject(&s3.PutObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(time.Now().Format("2006-01-02-15-04-05") + ".txt"),
				Body:   bytes.NewReader(msg.Value),
			})
			if err != nil {
				log.Printf("Failed to upload to S3: %s", err)
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
