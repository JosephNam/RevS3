package RevS3

import (
	"bytes"
	"fmt"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var service = s3.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))

func listFiles(bucket string) {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket), // Required
		/*
			ContinuationToken: aws.String("Token"),
			Delimiter:         aws.String("Delimiter"),
			EncodingType:      aws.String("EncodingType"),
			FetchOwner:        aws.Bool(true),
			MaxKeys:           aws.Int64(1),
			Prefix:            aws.String("Prefix"),
			RequestPayer:      aws.String("RequestPayer"),
			StartAfter:        aws.String("StartAfter"),
		*/
	}
	response, err := service.ListObjectsV2(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response)
}

func uploadFile(bucket string, file string) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bytesToSend := bytes.NewReader(dat)
	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(file),
		ACL:           aws.String("public-read"),
		Body:          bytesToSend,
		ContentLength: aws.Int64(int64(len(dat))),
		ContentType:   aws.String("mp3"),
		Metadata: map[string]*string{
			"Key": aws.String("MetadataValue"),
		},
	}
	service.PutObject(params)
}

func downloadFile(bucket string, file string) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(file + ".mp3"),
	}
	result, err := service.GetObject(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	fileErr := ioutil.WriteFile(file+"_new.mp3", buf.Bytes(), 0644)
	if fileErr != nil {
		fmt.Println(err.Error())
		return fileErr
	}
	fmt.Println("File write success")
	return fileErr
}
