package main

import (
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func uploadToS3Bucket(file io.Reader, fileName string) ImageUploadResponse {

	bucketName := os.Getenv("bucket_name")
	region := "us-east-2"

	//select Region to use.
	conf := aws.Config{Region: &region}
	sess, _ := session.NewSession(&conf)
	uploader := s3manager.NewUploader(sess)

	// Upload input parameters
	upParams := &s3manager.UploadInput{
		Bucket: &bucketName,
		Key:    &fileName,
		Body:   file,
	}

	// Perform an upload.
	result, err := uploader.Upload(upParams)

	if err != nil {
		log.Fatalf("Error in uploadToS3Bucket %v", err)
	}

	responseData := ImageUploadResponse{
		FileName: fileName,
		Location: result.Location,
	}

	return responseData
}
