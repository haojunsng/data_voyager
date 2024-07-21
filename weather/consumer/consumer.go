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
	common.HandleError(err, "Failed to create consumer")
	defer c.Close()

	err = c.Subscribe(common.KafkaTopic, nil)
	common.HandleError(err, "Failed to subscribe to topic")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(common.AWSRegion)},
	)
	common.HandleError(err, "Failed to create AWS Session")

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
		common.HandleError(err, "Failed to upload to S3")
	}
}

func uploadToS3(s3Svc *s3.S3, bucket string, message string) error {
	key := fmt.Sprintf("weather-data-%d.txt", time.Now().UnixNano())

	_, err := s3Svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(message),
	})

	common.HandleError(err, "Failed to put to S3")

	fmt.Printf("Successfully uploaded message to S3 with key %s\n", key)
	return nil
}
