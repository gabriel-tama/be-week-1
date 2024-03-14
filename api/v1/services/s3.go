package services

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Service interface {
	UploadFile(fileForm io.Reader) (*s3.PutObjectOutput, error)
}

type S3ServiceImpl struct {
	S3Client   *s3.Client
	bucketName string
	objectKey  string
}

func NewS3Service(accessKey string, secretKey string, bucketName string, objectKey string) S3Service {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))
	})

	return &S3ServiceImpl{
		S3Client:   client,
		bucketName: bucketName,
		objectKey:  objectKey,
	}
}

func (service *S3ServiceImpl) UploadFile(file io.Reader) (*s3.PutObjectOutput, error) {
	fileName := time.Now().String()
	bucketName := service.bucketName
	objectKey := "file.txt"

	obj, err := service.S3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		ACL:    types.ObjectCannedACL((*aws.String("public-read"))),
		Body:   file,
	})

	if err != nil {
		log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
			fileName, bucketName, objectKey, err)
	}

	return obj, err
}
