package main

import (
	"fmt"
	"strings"
	"os"
	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("arg err") // check number of cli arguments
		return
	}

	file := os.Args[1] // get file name

	fileinfo, err := os.Stat(file) 
	if err != nil {
		fmt.Println("file err") // check file exits
		return
	}
	size := fileinfo.Size()

	if size > 250000000 {
		fmt.Println("file too large") // check file size
		return
	}

	video, _ := gocv.VideoCaptureFile(file) // open file as video
	defer video.Close()

	var outfilename strings.Builder
	outfilename.WriteString(file[:strings.Index(file,".")])
	outfilename.WriteString("out")
	outfilename.WriteString(file[strings.Index(file,"."):])

	out, _ := gocv.VideoWriterFile(outfilename.String(),
								video.CodecString(),
								video.Get(gocv.VideoCaptureFPS),
								int(video.Get(gocv.VideoCaptureFrameWidth)),
								int(video.Get(gocv.VideoCaptureFrameHeight)),
								true)

	curr := gocv.NewMat() // reader mat
	defer curr.Close()

	for {

		if ok := video.Read(&curr); !ok {
			fmt.Println("video reading stopped") // read frame to reader mat
			return
		}

		if curr.Empty() { 
			continue
		}

		out.Write(curr)
	}
}