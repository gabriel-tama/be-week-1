package services

import (
	"context"
	"fmt"
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
	UploadFile(fileName string, fileForm io.Reader) (*s3.PutObjectOutput, error)
	GetObject(objectKey string) (*s3.GetObjectOutput, error)
	ListFile() (*s3.ListObjectsV2Output, error)
	GetObjectWithUrl(objectKey string) string
}

type S3ServiceImpl struct {
	S3Client   *s3.Client
	bucketName string
	baseUrl    string
	region     string
}

func NewS3Service(region string, accessKey string, secretKey string, bucketName string, baseUrl string) S3Service {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = region
		o.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))
	})

	return &S3ServiceImpl{
		S3Client:   client,
		bucketName: bucketName,
		baseUrl:    baseUrl,
	}
}

func (service *S3ServiceImpl) UploadFile(fileName string, file io.Reader) (*s3.PutObjectOutput, error) {
	bucketName := service.bucketName
	objectKey := fmt.Sprintf("%s/%v-%s", "ngab-gab", time.Now().Unix(), fileName)

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

func (service *S3ServiceImpl) ListFile() (*s3.ListObjectsV2Output, error) {
	bucketName := service.bucketName

	obj, err := service.S3Client.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		log.Printf("Couldn't list file on %s. Here's why: %v\n",
			bucketName, err)
	}

	return obj, err
}

func (service *S3ServiceImpl) GetObject(objectKey string) (*s3.GetObjectOutput, error) {
	bucketName := service.bucketName

	obj, err := service.S3Client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})

	if err != nil {
		log.Printf("Couldn't list file on %s. Here's why: %v\n",
			bucketName, err)
	}

	return obj, err
}

func (service *S3ServiceImpl) GetObjectWithUrl(objectKey string) string {
	return fmt.Sprintf("%s/%s", service.baseUrl, objectKey)
}
