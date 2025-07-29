package services

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type FileService interface {
	UploadToS3(ctx echo.Context, fileName string, buffer *bytes.Buffer) (string, error)
}

type S3Service struct {
	AccessKey    string
	SecretKey    string
	Region       string
	Bucket       string
	SessionToken string
}

func NewS3Service(
	accessKey string,
	secretKey string,
	region string,
	bucket string,
	sessionToken string,
) *S3Service {
	return &S3Service{
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		Region:       region,
		Bucket:       bucket,
		SessionToken: sessionToken,
	}
}

func (s *S3Service) UploadToS3(ctx echo.Context, fileName string, buffer *bytes.Buffer) (string, error) {
	awsCreds := credentials.NewStaticCredentialsProvider(
		s.AccessKey,
		s.SecretKey,
		s.SessionToken,
	)

	awsCfg, err := config.LoadDefaultConfig(
		ctx.Request().Context(),
		config.WithRegion(s.Region),
		config.WithCredentialsProvider(awsCreds),
	)
	if err != nil {
		return "", fmt.Errorf("Failed to load AWS config: %w", err)
	}

	s3Client := s3.NewFromConfig(awsCfg)
	id := uuid.New()
	pathName := fmt.Sprintf("ngikut/gogomanager/%s.png", &id)

	_, error := s3Client.PutObject(ctx.Request().Context(), &s3.PutObjectInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(pathName),
		Body:        buffer,
		ContentType: aws.String(http.DetectContentType(buffer.Bytes())),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if error != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.Bucket, s.Region, pathName)
	return fileURL, nil
}
