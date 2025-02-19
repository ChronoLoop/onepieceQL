package awsclient

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/cloudfront/sign"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client
var CloudFrontSigner *sign.URLSigner

const CLOUDFRONT_DOMAIN_NAME = "https://dok1t7q1m768g.cloudfront.net/"

func parseRSAPrivateKey(pemKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// Try parsing as PKCS8
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return key.(*rsa.PrivateKey), nil
	}

	return privKey, nil
}

func configS3() {
	AWS_REGION := os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(AWS_REGION))
	if err != nil {
		log.Fatal(err)
	}

	S3Client = s3.NewFromConfig(cfg)
}

func configCloudFront() {
	keyID := os.Getenv("CLOUDFRONT_KEY_PAIR_ID")
	privKeyEnv := os.Getenv("CLOUDFRONT_PRIVATE_KEY")
	privKey, err := parseRSAPrivateKey(privKeyEnv)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		os.Exit(1)
	}
	CloudFrontSigner = sign.NewURLSigner(keyID, privKey)
}

func ConfigAWS() {
	configS3()
	configCloudFront()
}
