package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	// Retrieve AWS credentials and region from environment variables
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	// Initialize AWS configuration with provided credentials and region
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion(region))
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		os.Exit(1)
	}

	// Create EC2 client
	client := ec2.NewFromConfig(cfg)

	// Retrieve EC2 instances
	result, err := client.DescribeInstances(context.Background(), &ec2.DescribeInstancesInput{})
	if err != nil {
		fmt.Println("Error describing instances:", err)
		os.Exit(1)
	}

	// Print instance details
	fmt.Printf("--------------: AWS Ec2 Instances in [%s] :-----------------\n", region)
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println("Instance ID:", *instance.InstanceId)
			fmt.Println("Instance State:", instance.State.Name)
			fmt.Println("Instance Type:", instance.InstanceType)
			fmt.Println("Instance Launch Time:", instance.LaunchTime)
			fmt.Println("Instance Private IP:", *instance.PrivateIpAddress)
			fmt.Println("Instance Public IP:", *instance.PublicIpAddress)
			fmt.Println("Instance Tags:")
			for _, tag := range instance.Tags {
				fmt.Printf("  %s: %s\n", *tag.Key, *tag.Value)
			}
			fmt.Println("-----------------------------------")
		}
	}
}
