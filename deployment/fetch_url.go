package main

import (
        "context"
      	"github.com/aws/aws-lambda-go/lambda"
)

type Evt struct {
        Fname string `json:"videofilename"`
}

func HandleRequest(ctx context.Context, event Evt) (string, error) {
        url := "https://thehotboxupload.s3.amazonaws.com/" + event.Fname

        rtrn := EditProcess(url)

        return rtrn, nil
}

func EditProcess(url string) (string) {
        return url
}

func main() {
        lambda.Start(HandleRequest)
}
