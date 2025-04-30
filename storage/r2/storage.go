package r2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	client     *s3.Client
	baseURL    string
	bucket     string
	baseFolder string
}

func Must(endpoint, accessKeyID, secretAccessToken, baseURL, bucket, baseFolder string) *Storage {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessToken, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(
		cfg,
		func(o *s3.Options) {
			o.BaseEndpoint = &endpoint
		},
	)

	if baseFolder != "" {
		baseFolder = fmt.Sprintf("%s/", baseFolder)
	}

	return &Storage{
		client:     client,
		baseURL:    baseURL,
		bucket:     bucket,
		baseFolder: baseFolder,
	}
}

func (s *Storage) Exists(ctx context.Context, path, name string) (exists bool, err error) {
	_, err = s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.pathname(path, name)),
	})
	if err != nil {
		var responseError *awshttp.ResponseError
		if errors.As(err, &responseError) && responseError.ResponseError.HTTPStatusCode() == http.StatusNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (s *Storage) Store(ctx context.Context, path, name string, data []byte) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.pathname(path, name)),
		Body:   bytes.NewReader(data),
	})

	return err
}

func (s *Storage) Load(ctx context.Context, path, name string) (data []byte, err error) {
	out, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.pathname(path, name)),
	})
	if err != nil {
		return
	}

	return io.ReadAll(out.Body)
}

func (s *Storage) Delete(ctx context.Context, path, name string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.pathname(path, name)),
	})

	return err
}

func (s *Storage) DeleteAll(ctx context.Context, path string) error {
	rsp, err := s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(path),
	})
	if err != nil {
		return err
	}

	for _, obj := range rsp.Contents {
		_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    obj.Key,
		})
		if err != nil {
			return err
		}
	}

	for *rsp.IsTruncated {
		rsp, err = s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket:            aws.String(s.bucket),
			Prefix:            aws.String(path),
			ContinuationToken: rsp.NextContinuationToken,
		})
		if err != nil {
			return err
		}

		for _, obj := range rsp.Contents {
			_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
				Bucket: aws.String(s.bucket),
				Key:    obj.Key,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Storage) Copy(ctx context.Context, srcPath, srcName, dstPath, dstName string) error {
	_, err := s.client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(s.bucket),
		CopySource: aws.String(fmt.Sprintf("%s/%s", s.bucket, s.pathname(srcPath, srcName))),
		Key:        aws.String(s.pathname(dstPath, dstName)),
	})

	return err
}

func (s *Storage) GenerateSignedURL(path string, name string) (url string, err error) {
	pc := s3.NewPresignClient(s.client)

	req, err := pc.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.pathname(path, name)),
	}, s3.WithPresignExpires(15*time.Minute))
	if err != nil {
		return "", err
	}

	return req.URL, nil
}

func (s *Storage) pathname(path string, name string) string {
	o := ""

	if len(path) > 0 {
		o = fmt.Sprintf("%s%s/%s", s.baseFolder, path, name)
	} else {
		o = name
	}

	return o
}
