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
	fmt.Printf("Enter Region: ")
	fmt.Scanln(&region)
	fmt.Printf("Enter Bucket Name: ")
	fmt.Scanln(&myBucket)
	fmt.Printf("Object Name(Key):")
	fmt.Scanln(&myKey)
	
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := s3.New(sess, &aws.Config{Region: aws.String(region)})

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myKey),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	//log.Println("The URL is", urlStr)
	fmt.Printf("The URL is %s", urlStr)
}
