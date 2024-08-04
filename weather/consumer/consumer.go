package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"weather/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/xitongsys/parquet-go-source/writerfile"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
)

type Consumer struct {
	Consumer *kafka.Consumer
	S3Client *s3.S3
	Bucket   string
}

func NewConsumer() (*Consumer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": common.KafkaBroker,
		"group.id":          common.GroupID,
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	err = consumer.Subscribe(common.KafkaTopic, nil)
	if err != nil {
		return nil, err
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(common.AWSRegion),
	})
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(sess)

	return &Consumer{
		Consumer: consumer,
		S3Client: s3Client,
		Bucket:   common.S3Bucket,
	}, nil
}

func (c *Consumer) PollMessages() {
	for {
		ev := c.Consumer.Poll(int(100 * time.Millisecond))

		switch msg := ev.(type) {
		case *kafka.Message:
			messageValue := string(msg.Value)
			fmt.Printf("received message: %s\n", messageValue)

			offset := msg.TopicPartition.Offset
			fmt.Printf("message offset: %v\n", offset)

			var weatherData common.WeatherData
			// Unmarshal JSON data into struct to quickly validate data
			err := json.Unmarshal(msg.Value, &weatherData)
			if err != nil {
				log.Printf("failed to unmarshal message: %v\n", err)
				continue
			}

			err = c.UploadToS3(weatherData, offset)
			if err != nil {
				log.Printf("failed to upload message to S3: %v", err)
			}
		case kafka.Error:
			if msg.IsFatal() {
				log.Fatalf("fatal Kafka error: %v", msg)
			} else {
				log.Printf("kafka error: %v", msg)
			}
		default:
			log.Printf("no new events to be polled...")
		}
	}
}

func (c *Consumer) UploadToS3(data common.WeatherData, offset kafka.Offset) error {
	date := time.Now().Format("20060102")
	key := fmt.Sprintf("source=kafka/type=weather/date=%s/%d-%d.parquet", date, time.Now().Unix(), offset)

	var buf bytes.Buffer
	fw := writerfile.NewWriterFile(&buf)
	pw, err := writer.NewParquetWriter(fw, new(common.WeatherData), 4)
	if err != nil {
		return fmt.Errorf("failed to create parquet writer: %w", err)
	}

	pw.RowGroupSize = 128 * 1024 * 1024
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	fmt.Printf("Data to write: %+v\n", data)

	if err = pw.Write(data); err != nil {
		return fmt.Errorf("failed to write data to parquet: %w", err)
	}

	if err = pw.WriteStop(); err != nil {
		return fmt.Errorf("failed to stop parquet writer: %w", err)
	}

	if err = fw.Close(); err != nil {
		return fmt.Errorf("failed to close parquet file writer: %w", err)
	}

	fmt.Printf("Buffer size before upload: %d bytes\n", buf.Len())
	fmt.Printf("Buffer contents: %v\n", buf.Bytes())

	_, err = c.S3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return fmt.Errorf("failed to upload message to S3: %w", err)
	}

	fmt.Printf("Uploaded message to S3 with key: %s\n", key)
	return nil
}
