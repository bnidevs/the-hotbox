package main

import (
    "context"
    "log"
    "github.com/aws/aws-lambda-go/lambda"
	  "gocv.io/x/gocv"
)

type Evt struct {
    Fname string `json:"videofilename"`
}

func HandleRequest(ctx context.Context, event Evt) (gocv.Mat, error) {
    url := "https://thehotboxupload.s3.amazonaws.com/" + event.Fname

    vid, err := gocv.VideoCaptureFile(url)
    if(err != nil){
        log.Fatal(err)
    }
    defer vid.Close()

    rtrn := gocv.NewMat()
    defer rtrn.Close()

    vidcap := gocv.VideoCapture(*vid)
    readstat := vidcap.Read(&rtrn)

    if(!readstat){
        log.Fatal(readstat)
    }

    return rtrn, nil
}

func main() {
    lambda.Start(HandleRequest)
}
