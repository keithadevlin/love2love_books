//go:generate mockgen -package service -source=service.go -destination service_mocks.go

package service

import (
	"context"
	"github.com/sirupsen/logrus"
)

type PersistFile interface {
	PersistFilePayload(ctx context.Context, fileName string, payload string) (string, error)
}

type Service struct {
	cfg         *Config
	persistFile PersistFile
}

type Config struct {
	SalesReportBucketName string
}

func NewService(cfg *Config, persistFile PersistFile) *Service {
	return &Service{
		cfg:         cfg,
		persistFile: persistFile,
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
