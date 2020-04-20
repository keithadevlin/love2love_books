//go:generate mockgen -package handler -source=handler.go -destination handler_mock.go

package handler

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/keithadevlin/love2love_books/pkg/shared/mws"
	"io"
	"log"
	"os"
)

//Service defines the Service interface.
type Service interface {
	CreateInvoice(ctx context.Context, orderReportItem mws.OrderReportItem, invoiceNumber int) error
}

type Handler struct {
	service   Service
	inputFile *os.File
	log       log.Logger
}

func NewHandler(s Service, log log.Logger, inputFile *os.File) *Handler {
	return &Handler{
		service:   s,
		log:       log,
		inputFile: inputFile,
	}

}

func (h Handler) OrderInvoiceGeneration(ctx context.Context, startInvoiceNumber int) error {
	// Parse the file
	r := csv.NewReader(h.inputFile)
	r.Comma = '\t'
	invoiceNumber := startInvoiceNumber

	lineCount := 0
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println("Next Invoice run should start at ", invoiceNumber)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//need this to ignore the header line
		lineCount ++
		if lineCount == 1 {
			continue
		}
		orderReportItem := mws.OrderReportItem{
			OrderId:              record[0],
			OrderItemId:          record[1],
			PurchaseDate:         record[2],
			PaymentsDate:         record[3],
			BuyerEmail:           record[4],
			BuyerName:            record[5],
			BuyerPhoneNumber:     record[6],
			Sku:                  record[7],
			ProductName:          record[8],
			QuantityPurchased:    record[9],
			Currency:             record[10],
			ItemPrice:            record[11],
			ItemTax:              record[12],
			ShippingPrice:        record[13],
			ShippingTax:          record[14],
			ShipServiceLevel:     record[15],
			RecipientName:        record[16],
			ShipAddress1:         record[17],
			ShipAddress2:         record[18],
			ShipAddress3:         record[19],
			ShipCity:             record[20],
			ShipState:            record[21],
			ShipPostalCode:       record[22],
			ShipCountry:          record[23],
			ShipPhoneNumber:      record[24],
			DeliveryStartDate:    record[24],
			DeliveryEndDate:      record[25],
			DeliveryTimeZone:     record[26],
			DeliveryInstructions: record[27],
			SalesChannel:         record[28],
		}


		fmt.Printf("Processing Invoice for order %s\n", record[0])

		err = h.service.CreateInvoice(ctx, orderReportItem, invoiceNumber)
		if err != nil {
			return err
		}
		invoiceNumber++

	}
	return nil
}
