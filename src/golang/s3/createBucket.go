package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"time"
)

var bucket = "bucket-o-fun"

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO() /*, config.WithSharedConfigProfile("my-profile"), config.WithRegion("us-east-2")*/)

	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	result, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CreateBucket Location: %s", result.Location)

	waiter := s3.NewBucketExistsWaiter(client)
	maxWaitTime := time.Second * 30

	err = waiter.Wait(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	}, maxWaitTime)

	if err != nil {
		log.Fatal(err)
	}
}
