package video

import (
	"fmt"
	"sync"
	"gocv.io/x/gocv"
	"../image"
)

var mutex = &sync.Mutex{}
var totalEditor int = 4

func editFrame(frameEditor chan bool, frames []*gocv.Mat, editorNum int, 
				videoIn *gocv.VideoCapture, totalFrames int, change int16) {
	start := totalFrames*editorNum/totalEditor
	end := totalFrames*(editorNum+1)/totalEditor
	for i:=start; i<end; i++ {
		// curr := gocv.NewMat()
		// frames[i] = &curr
		
		//Trying to implement synchronous read here

		// //fmt.Println("hdwuhdw",videoIn.Get(gocv.VideoCapturePosFrames))
		// mutex.Lock()
		// videoIn.Set(gocv.VideoCapturePosFrames, float64(i))
		// //fmt.Println(videoIn.Get(gocv.VideoCapturePosFrames))
		// if ok := videoIn.Read(&curr); !ok {
		// 	fmt.Println("failed",i,totalFrames)
		// 	return
		// }
		// mutex.Unlock()

		image.ModifyBrightness(frames[i], change)
		if frames[i].Empty() {
			fmt.Println("empty")
		}
		//fmt.Println(i, end, totalFrames, editorNum)
	}
	frameEditor <- true


}

func ModifyBrightnessSync(videoIn *gocv.VideoCapture, videoOut *gocv.VideoWriter, change int16) {
	totalFrames := int(videoIn.Get(gocv.VideoCaptureFrameCount))

	frameEditors := make([]chan bool, totalEditor)
	frames := make([]*gocv.Mat, totalFrames)

	//sequential read
	for i:=0; i<totalFrames; i++ {
		curr:= gocv.NewMat()
		frames[i] = &curr
		videoIn.Read(frames[i])
	}

	//create 4 channels to process frames
	for i:=0; i<totalEditor; i++ {
		frameEditors[i] = make(chan bool)
		go editFrame(frameEditors[i], frames, i, videoIn, totalFrames, change)
	}


	for i:=0; i<len(frameEditors); i++ {
		<-frameEditors[i]
		
		start := totalFrames*i/totalEditor
		end := totalFrames*(i+1)/totalEditor
		for j:=start; j<end; j++ {
			videoOut.Write(*frames[j])
		}
		
		close(frameEditors[i])
	}

}