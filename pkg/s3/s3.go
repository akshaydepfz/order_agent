package s3

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Config struct {
	Bucket  string
	Region  string
	Key     string
	Secret  string
}

type Client struct {
	client *s3.Client
	cfg    Config
}

func NewClient(cfg Config) (*Client, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.Key, cfg.Secret, "")),
	)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: s3.NewFromConfig(awsCfg),
		cfg:    cfg,
	}, nil
}

func (c *Client) UploadImage(ctx context.Context, key string, body []byte, contentType string) (string, error) {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(c.cfg.Bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(body),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", c.cfg.Bucket, c.cfg.Region, key), nil
}

func (c *Client) GenerateShopLogoKey(filename string) string {
	return fmt.Sprintf("ShopLogos/%d_%s", time.Now().UnixNano(), filename)
}
