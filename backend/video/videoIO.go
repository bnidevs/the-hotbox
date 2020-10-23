package video

import (
	"fmt"
	"strings"
	"os"
	"gocv.io/x/gocv"
)


func OpenVideo(fileName string) *gocv.VideoCapture {
	fileinfo, err := os.Stat(fileName) 
	if err != nil {
		fmt.Println("file err") // check file exits
		os.Exit(1)
		//return nil
	}
	size := fileinfo.Size()

	if size > 250000000 {
		fmt.Println("file too large") // check file size
		os.Exit(1)
	}

	videoCaptured, _ := gocv.VideoCaptureFile(fileName) // open file as video
	return videoCaptured
}

//Name the output file of video
//Parameter: fileName  (recommend the name of the input video file)
//Return: string
func NameOut(fileName string) string {
	var outFileName strings.Builder
	outFileName.WriteString(fileName[:strings.Index(fileName,".")])
	outFileName.WriteString("out")
	outFileName.WriteString(fileName[strings.Index(fileName,"."):])
	return outFileName.String()
}
