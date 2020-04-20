package main

import (
	"context"
	"github.com/keithadevlin/love2love_books/pkg/lambda/create_amazon_invoice/internal/handler"
	"github.com/keithadevlin/love2love_books/pkg/lambda/create_amazon_invoice/internal/service"
	"log"
	"os"
)

func main() {

	//SET This VALUE PRIOR TO EACH RUN
	startInvoiceNumber := 2000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	log := log.Logger{}

	s := service.NewService(log)

	csvfile, err := os.Open("/Users/keith/go/src/github.com/keithadevlin/love2love_books/pkg/lambda/create_amazon_invoice/data/in/orders.tsv")
	if err != nil {
		log.Fatalf("Couldn't open the csv file", err)
	}
	defer csvfile.Close()

	h := handler.NewHandler(s, log, csvfile)

	err = h.OrderInvoiceGeneration(ctx, startInvoiceNumber)

	if err != nil {
		log.Fatalf("Couldn't generate invoice", err)
	}

}
