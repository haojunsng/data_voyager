package main

import (
	"fmt"
	"strings"
	"time"
	"weather/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consume() error {
	config := kafka.ConfigMap{
		"bootstrap.servers": common.KafkaBroker,
		"group.id":          common.GroupID,
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(&config)
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
	defer c.Close()

	err = c.Subscribe(common.KafkaTopic, nil)
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(common.AWSRegion)},
	)
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}

	s3Svc := s3.New(sess)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Code() == kafka.ErrTimedOut {
				// TODO : Handle error
				continue
			}
			return fmt.Errorf("failed to read message: %w", err)
		}

		fmt.Printf("Received message: %s\n", string(msg.Value))

		err = uploadToS3(s3Svc, common.S3Bucket, string(msg.Value))
		if err != nil {
			log.Fatalf("%s: %s", message, err)
		}
	}
}

func uploadToS3(s3Svc *s3.S3, bucket string, message string) error {
	key := fmt.Sprintf("weather-data-%d.txt", time.Now().UnixNano())

	_, err := s3Svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(message),
	})

	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}

	fmt.Printf("Successfully uploaded message to S3 with key %s\n", key)
	return nil
}
