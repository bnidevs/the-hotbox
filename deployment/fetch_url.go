package main

import (
        "context"
      	"github.com/aws/aws-lambda-go/lambda"
        "github.com/aws/aws-sdk-go/aws"
)

type Evt struct {
        Fname string `json:"videofilename"`
}

func HandleRequest(ctx context.Context, event Evt) (string, error) {
        return "https://thehotboxupload.s3.amazonaws.com/" + event.Fname
}

func main() {
        lambda.Start(HandleRequest)
}
