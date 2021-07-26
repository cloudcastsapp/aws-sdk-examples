package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
	"time"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO() /*, config.WithSharedConfigProfile("cloudcasts"), config.WithRegion("us-east-2")*/)

	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := ec2.NewFromConfig(cfg)

	result, err := client.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		// ImageId: https://cloud-images.ubuntu.com/locator/ec2/
		ImageId: aws.String("ami-0b29b6e62f2343b46"),
		InstanceType: types.InstanceTypeT3Small,

		// Optional parameters you likely want to define
		// KeyName: aws.String("some-key"),
		// SecurityGroupIds: []string{
		//	"sg-foobar",
		// },
		// SubnetId: aws.String("subnet-foobar"),

		MaxCount: aws.Int32(1),
		MinCount: aws.Int32(1),
		BlockDeviceMappings: []types.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sda1"),
				Ebs: &types.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true),
					VolumeType: "gp3",
					VolumeSize: aws.Int32(8),
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Instance ID: %s", result.Instances[0].InstanceId)

	waiter := ec2.NewInstanceRunningWaiter(client)
	maxWaitTime := time.Second * 30

	err = waiter.Wait(context.TODO(), &ec2.DescribeInstancesInput{
		InstanceIds: []string {
			*result.Instances[0].InstanceId,
		},
	}, maxWaitTime)

	if err != nil {
		log.Fatal(err)
	}
}
