//go:generate mockgen -package handler -source=handler.go -destination handler_mock.go

package handler

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Service interface {
	WriteSalesReport(ctx context.Context) error
	RequestMWSReport(ctx context.Context, reportName string) (string, error)
}

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}

}

func (h *Handler) ProcessSalesReport(ctx context.Context) error {

	logrus.Infof("Requesting Report")
	reportRequestId, err := h.service.RequestMWSReport(ctx, "_GET_MERCHANT_LISTINGS_DATA_BACK_COMPAT_")
	if err != nil {
		logrus.Infof("error in process sales report ", err)
		return err
	}
	logrus.Infof("ReportRequstId = %s", reportRequestId)

	logrus.Infof("in processSales Report")
	err = h.service.WriteSalesReport(ctx)
	if err != nil {
		logrus.Infof("error in process sales report ", err)
		return err
	}
	logrus.Infof("success in process sales report ")
	return nil
}
