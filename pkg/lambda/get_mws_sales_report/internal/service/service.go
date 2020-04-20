//go:generate mockgen -package service -source=service.go -destination service_mocks.go

package service

import (
	"context"
	"encoding/xml"
	"github.com/keithadevlin/love2love_books/pkg/shared/mws"
	"github.com/sirupsen/logrus"
)

type PersistFile interface {
	PersistFilePayload(ctx context.Context, fileName string, payload string) (string, error)
}

type Service struct {
	cfg         *Config
	persistFile PersistFile
	mwsapi      mws.AmazonMWSAPI
}

type Config struct {
	SalesReportBucketName string
}

func NewService(cfg *Config, persistFile PersistFile, mwsapi mws.AmazonMWSAPI) *Service {
	return &Service{
		cfg:         cfg,
		persistFile: persistFile,
		mwsapi:      mwsapi,
	}

}

func (s *Service) WriteSalesReport(ctx context.Context) error {
	logrus.Infof("in writeSales Report")
	filename, err := s.persistFile.PersistFilePayload(ctx, "test", "test123")
	if err != nil {
		logrus.Infof("error writing file ", err)
		return err
	}
	logrus.Infof("filename = ", filename)
	return nil
}

func (s *Service) RequestMWSReport(ctx context.Context, reportName string) (string, error) {

	logrus.Info("Submitting report request")

	result, err := s.mwsapi.RequestReport(reportName)

	//result1, err5 := api2.RequestReport("_GET_MERCHANT_LISTINGS_DATA_BACK_COMPAT_")

	logrus.Info("Submitted report request")

	if err != nil {
		return "", err
	}

	logrus.Info(result)

	byteResult1 := []byte(result)
	var reportRequestRequestReportResponse mws.ReportRequestRequestReportResponse
	err = xml.Unmarshal(byteResult1, &reportRequestRequestReportResponse)
	if err != nil {
		logrus.Info("Error unmarshalling report request")
		return "", err
	}

	reportRequestId := string(reportRequestRequestReportResponse.ReportRequestRequestReportResult.ReportRequestReportRequestInfo.ReportRequestReportRequestId.Id)
	logrus.Info("Report Request ID as string = %s", reportRequestId)

	return reportRequestId, nil
}
