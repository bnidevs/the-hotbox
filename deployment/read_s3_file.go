package main

import (
        "fmt"
        "io/ioutil"
        "context"
        "log"
      	"github.com/aws/aws-lambda-go/events"
      	"github.com/aws/aws-lambda-go/lambda"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        s3session := s3.New(session.New())

        getObjInput := &s3.GetObjectInput{Bucket: aws.String("thehotboxupload"), Key: aws.String(request.QueryStringParameters["videofilename"])}

        result, err := s3session.GetObject(getObjInput)
        if(err != nil){
            log.Fatal(err)
        }

        blob, err := ioutil.ReadAll(result.Body)
        if(err != nil){
            log.Fatal(err)
        }

        defer result.Body.Close()

        rtrnstr := fmt.Sprintf("%s", blob)

        rtrn := events.APIGatewayProxyResponse{StatusCode: 200, Body: rtrnstr}

        return rtrn, nil
}

func main() {
        lambda.Start(HandleRequest)
}
