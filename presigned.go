package main

import (
	"fmt"
	"time"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

)

var(
	myBucket string
	region string
	myKey string
)

func main() {
	// console enter region, bucket, object
	fmt.Printf("Enter Region: ")
	fmt.Scanln(&region)
	fmt.Printf("Enter Bucket Name: ")
	fmt.Scanln(&myBucket)
	fmt.Printf("Object Name(Key):")
	fmt.Scanln(&myKey)
	
	// create aws session
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}
	// create s3 connection
	svc := s3.New(sess, &aws.Config{Region: aws.String(region)})

	// create a s3.getobjectrequest using local s3 sdk
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myKey),
	})

	// Sign url using local credential and set timeout = 15 mins
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		log.Println("Failed to sign request", err)
	}

	// Print URL at console 
	fmt.Printf("The URL is %s", urlStr)
}
