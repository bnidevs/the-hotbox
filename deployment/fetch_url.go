package main

import (
        "context"
      	"github.com/aws/aws-lambda-go/lambda"
)

type Evt struct {
        Fname string `json:"videofilename"`
        Brightness int `json:"brightness_val"`
        Saturation int `json:"saturation_val"`
        Contrast int `json:"contrast_val"`
        Noise int `json:"noise_val"`
        LaserEyes bool `json:"laser_eyes_check"`
        HeadBulge bool `json:"head_bulge_check"`
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
