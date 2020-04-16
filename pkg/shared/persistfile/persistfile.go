//go:generate mockgen -package persistfile -source=persistfile.go -destination persistfile_mocks.go

package persistfile

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)




//type S3Client interface {
//	s3iface.S3API
//}

//Client
type Client struct {
	BucketName   string
	FunctionName string
	Uploader     *s3manager.Uploader
	//s3cli      S3Client
}

//NewPersistFailedPayloadClient
func NewClient(bucketName string, functionName string, uploader *s3manager.Uploader) *Client {
	return &Client{
		BucketName:   bucketName,
		FunctionName: functionName,
		Uploader:     uploader,
	}
}

//PersistFailedPayload
func (c *Client) PersistFilePayload(ctx context.Context, fileName string, payload string) (string, error) {

	key := fmt.Sprintf("%s/%s-%s.csv", c.FunctionName, fileName, time.Now().Format("20060102_150405"))

	// create a reader from data data in memory
	reader := strings.NewReader(payload)

	_, err := c.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.BucketName),
		Key:    aws.String(key),
		Body: reader,
	})
	if err != nil {
		logrus.Infof("error storing errored payload to failure bucket. Function %s, Filename: %s, Bucket: %s : %s", c.FunctionName, fileName, c.BucketName, err)
		fmt.Println(err)
		return key, fmt.Errorf("error storing errored payload to failure bucket. Function %s, Filename: %s, Bucket: %s", c.FunctionName, fileName, c.BucketName)

	}

	fmt.Printf("Successfully uploaded %q to %q\n", key, c.BucketName)
	return key, nil
}
