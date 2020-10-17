package main

import (
        "context"
        "log"
      	"github.com/aws/aws-lambda-go/lambda"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/aws"
)

type Evt struct {
        Fname string `json:"videofilename"`
}

func HandleRequest(ctx context.Context, event Evt) (string, error) {
        s3session := s3.New(session.New())

        getObjInput := &s3.GetObjectInput{Bucket: aws.String("thehotboxupload"), Key: aws.String(event.Fname)}

        result, err := s3session.GetObject(getObjInput)
        if(err != nil){
            log.Fatal(err)
        }

        rtrn := "upload good"

        if(*result.ContentLength > 250000000){
            delObjInput := &s3.DeleteObjectInput{Bucket: aws.String("thehotboxupload"), Key: aws.String(event.Fname)}
            _, err := s3session.DeleteObject(delObjInput)
            if(err != nil){
                log.Fatal(err)
            }
            rtrn = "file too big"
        }

        return rtrn, nil
}

func main() {
        lambda.Start(HandleRequest)
}
