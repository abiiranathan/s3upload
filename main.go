package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

// Define flags
var (
	fileName   = ""
	bucketName = ""
	objectKey  = ""

	// Authentication
	accessKey = ""
	secretKey = ""
	region    = "eu-west-2"
)

func init() {
	flag.StringVar(&fileName, "file", "", "File to upload to S3")
	flag.StringVar(&bucketName, "bucket", "", "S3 bucket name")
	flag.StringVar(&objectKey, "key", "", "Object key in S3")

	flag.StringVar(&accessKey, "access-key", "", "AWS access key")
	flag.StringVar(&secretKey, "secret-key", "", "AWS secret key")
	flag.StringVar(&region, "region", region, "AWS region")

	godotenv.Load()

}

func main() {

	// Parse command-line arguments
	flag.Parse()

	// Check if required flags are provided
	if fileName == "" || bucketName == "" || objectKey == "" {
		fmt.Println("Usage: go run main.go -file <file_path> -bucket <bucket_name> -key <object_key> -access-key <accessKey (Optional)> -secret-key <secret key(Optional)>")
		os.Exit(1)
	}

	if accessKey == "" {
		accessKey = os.Getenv("AWS_ACCESS_KEY")
	}

	if secretKey == "" {
		secretKey = os.Getenv("AWS_SECRET_KEY")
	}

	// Create an S3 session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}))

	// Create an S3 client
	s3Client := s3.New(sess, aws.NewConfig().WithRegion(region))

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Upload the file to S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   aws.ReadSeekCloser(file),
	})

	if err != nil {
		fmt.Println("Error uploading file to S3:", err)
		os.Exit(1)
	}
	fmt.Printf("File '%s' uploaded to S3 bucket '%s' with key '%s'\n", fileName, bucketName, objectKey)
}
