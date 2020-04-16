package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/keithadevlin/love2love_books/pkg/lambda/get_mws_sales_report/internal/handler"
	"github.com/keithadevlin/love2love_books/pkg/lambda/get_mws_sales_report/internal/service"
	"github.com/keithadevlin/love2love_books/pkg/shared/aws/awscore"
	"github.com/keithadevlin/love2love_books/pkg/shared/persistfile"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type environment struct {
	SalesReportBucketName string `envconfig:"SALES_REPORT_BUCKET_NAME" required:"true"`
	AWSRegion             string `envconfig:"REGION" required:"true"`
}

const (
	serviceName = "get-mws-sales-report"
)

func main() {

	var config environment
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal("error processing environment configuration")
	}

	cfg := &service.Config{
		SalesReportBucketName: config.SalesReportBucketName,
	}

	sess, err := awscore.NewSessionWithRegion(config.AWSRegion)
	if err != nil {
		log.Fatalf("error creating AWS Session", err)
	}

	//log.Infof("creating  lambda store")
	//store := awss3.NewBlockstore(sess)
	log.Infof("creating  lambda uploaders")
	uploader := s3manager.NewUploader(sess)
	log.Infof("finished creating lambda uploader")
	persistFileClient := persistfile.NewClient(config.SalesReportBucketName, serviceName, uploader)

	s := service.NewService(cfg, persistFileClient)
	log.Infof("finished creating new service")
	h := handler.NewHandler(s)
	log.Infof("finished creating new handler")

	log.Infof("staring Lambda")
	lambda.Start(h.ProcessSalesReport)
	log.Infof("finished executing lambda Lambda")
}
