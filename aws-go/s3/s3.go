package main

import (
	"fmt"

	aws "github.com/aws/aws-sdk-go/aws"
	session "github.com/aws/aws-sdk-go/aws/session"
	s3 "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		fmt.Errorf(err.Error())
	}
	s3c := s3.New(sess)
	input := &s3.ListBucketsInput{}
	buckets, err := s3c.ListBuckets(input)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(buckets)
	export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
	export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
	export AWS_DEFAULT_REGION=us-west-2
}
