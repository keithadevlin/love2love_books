//go:generate mockgen -package service -source=service.go -destination service_mocks.go

package service

import (
	"context"
	"fmt"
	"github.com/angelodlfrtr/go-invoice-generator"
	"github.com/keithadevlin/love2love_books/pkg/shared/mws"
	"io/ioutil"
	"log"
	"strconv"
)

type Invoice struct {
	//generator.New(docType string, options *Options) (*Document, error)

}

type Service struct {
	logger   log.Logger
}

func NewService(logger log.Logger) *Service {
	return &Service{
		logger:   logger,
	}
}

func (s Service) CreateInvoice(ctx context.Context, orderReportItem mws.OrderReportItem, invoiceNumber int) error {

	document, err := generator.New(generator.Invoice, &generator.Options{
		TextTypeInvoice: "Invoice",
		AutoPrint:       true,
		TextDateTitle: "Purchase Date:",
		CurrencySymbol: "£",
		CurrencyDecimal: ".",
		TextRefTitle: "Invoice Number",
		TextItemsTaxTitle: "VAT",
		TextItemsTotalHTTitle: "Net",
		TextItemsTotalTTCTitle: "Gross",
		TextTotalTax: "VAT",
		TextTotalWithTax: "Gross",
		TextTotalTotal: "Net",
	})

	if err != nil {
		return err
	}

	headerText := fmt.Sprintf("<center>Amazon Order: %s </center>", orderReportItem.OrderId)
	document.SetHeader(&generator.HeaderFooter{
		Text:       headerText,
		Pagination: true,
	})

	document.SetFooter(&generator.HeaderFooter{
		Text:       "<center>Many thanks for your custom</center>",
		Pagination: true,
	})

	document.SetRef(strconv.Itoa(invoiceNumber))
	//doc.SetVersion("someversion")
	shortPurchaseDate := orderReportItem.PurchaseDate[0:10]
	document.SetDate(shortPurchaseDate)

	document.SetDescription("Items")

	logoBytes, _ := ioutil.ReadFile("/Users/keith/go/src/github.com/keithadevlin/love2love_books/pkg/lambda/create_amazon_invoice/cmd/newLogo.png")

	document.SetCompany(&generator.Contact{
		Name: "Love2Love Books",
		Logo: &logoBytes,
		Address: &generator.Address{
			Address:    "7 Maudsville Cottages",
			Address2:   "Hanwell",
			PostalCode: "W7 3TE",
			City:       "London",
			Country:    "UK",
		},
	})

	document.SetCustomer(&generator.Contact{
		Name: orderReportItem.BuyerName,
		Address: &generator.Address{
			Address:    orderReportItem.ShipAddress1,
			Address2:   orderReportItem.ShipAddress2,
			PostalCode:       orderReportItem.ShipCity,
			City: orderReportItem.ShipPostalCode,
			Country:    orderReportItem.ShipCountry,
		},
	})

	shortProductName := orderReportItem.ProductName
	if len(orderReportItem.ProductName) > 55 {
		shortProductName = orderReportItem.ProductName[0:55]
	}

	for i := 0; i < 1; i++ {
		document.AppendItem(&generator.Item{
			Name:     shortProductName,
			UnitCost: orderReportItem.ItemPrice,
			Quantity: orderReportItem.QuantityPurchased,
			Tax: &generator.Tax{
				Percent: "0",
			},
		})
	}

	pdf, err := document.Build()

	if err != nil {
		return err
	}

	outputFileName := fmt.Sprintf("/Users/keith/go/src/github.com/keithadevlin/love2love_books/pkg/lambda/create_amazon_invoice/data/out/%s-#%s-%s.pdf", strconv.Itoa(invoiceNumber), orderReportItem.OrderId, shortProductName)

	err = pdf.OutputFileAndClose(outputFileName)

	if err != nil {
		return err
	}

	return nil
}
