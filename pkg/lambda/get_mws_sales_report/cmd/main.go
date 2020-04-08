package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	log.Printf("Executing Lambda")
	return fmt.Sprintf("Hello world"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
