package video

import (
	"container/heap"
	"fmt"
	"sync"

	"../image"
	"../utils"
	"gocv.io/x/gocv"
)

const threadLimit int = 5

func ModifyVideoSequential(videoIn *gocv.VideoCapture, videoOut *gocv.VideoWriter, paramlist utils.Parameters) {
	curr := gocv.NewMat()
	defer curr.Close()

	var faceClassifier gocv.CascadeClassifier
	var eyeClassifier gocv.CascadeClassifier
	if paramlist.LaserEye {
		faceXmlLoc := "image/detection/haarcascades/haarcascade_frontalface_default.xml"
		faceClassifier = gocv.NewCascadeClassifier()
		faceClassifier.Load(faceXmlLoc)
		defer faceClassifier.Close()

		eyeXmlLoc:= "image/detection/haarcascades/haarcascade_eye.xml"
		eyeClassifier = gocv.NewCascadeClassifier()
		eyeClassifier.Load(eyeXmlLoc)
		defer eyeClassifier.Close()
	}

	paramlist.CurrFrame = 0
	for ok := videoIn.Read(&curr); ok; ok = videoIn.Read(&curr) {
		if curr.Empty() {
			continue
		}
		paramlist.CurrFrame += 1

		if paramlist.LaserEye {
			image.LaserEyes(&curr, &eyeClassifier, &faceClassifier)
		}
		
		image.ModifyAll(&curr, paramlist)
		videoOut.Write(curr)
	}
}

func ModifyVideoThreaded(videoIn *gocv.VideoCapture, videoOut *gocv.VideoWriter, paramlist utils.Parameters) {
	curr := gocv.NewMat() // reader mat
	defer curr.Close()

	var wg sync.WaitGroup
	var waitPush sync.WaitGroup
	pq := make(utils.PriorityQueue, 0)

	// goes through the frames and modifies them, threaded
	totalFrames := int(videoIn.Get(gocv.VideoCaptureFrameCount))
	for i := 0; i < totalFrames; i += threadLimit {

		// Limit the number of active threads at any given time
		for j := 0; j < threadLimit; j++ {
			ok := videoIn.Read(&curr)

			if !ok {
				break
			} else if curr.Empty() {
				continue
			}

			wg.Add(1)

			go func(j int) {
				defer wg.Done()
				defer waitPush.Done()

				image.ModifyAll(&curr, paramlist)

				frame := &utils.Frame{
					Image:    curr,
					Priority: j,
				}

				waitPush.Wait()
				waitPush.Add(1)
				heap.Push(&pq, frame)
			}(j)

			// Blocking here makes it behave sequentially
			// wg.Wait()
		}

		// Blocking here seems to fail
		wg.Wait()

		if pq.Len() < threadLimit {
			fmt.Println("Possible Discrepancy")
		}

		for pq.Len() > 0 {
			frame := heap.Pop(&pq).(*utils.Frame)
			videoOut.Write(frame.Image)
		}
	}
}

/*
func ModifyBrightness(videoIn *gocv.VideoCapture, videoOut *gocv.VideoWriter, change int16) {
	curr := gocv.NewMat() // reader mat
	defer curr.Close()

	for {
		// the frame curr is of type CV8UC3
		// CV8U means it stores unsigned chars (lit. 8 bit Unsigned)
		// C3 means it has 3 channels, these channels represent the Blue Green and Red respectively
		if ok := videoIn.Read(&curr); !ok {
			return
		}

		if curr.Empty() {
			continue
		}

		image.ModifyBrightness(&curr, change)
		videoOut.Write(curr)

	}
}

func ModifyContrast(videoIn *gocv.VideoCapture, videoOut *gocv.VideoWriter, alpha float64) {
	ModifyVideo(image.ModifyContrast, videoIn, videoOut, alpha)
}

func ModifySaturation(videoIn *gocv.VideoCapture, videoOut *gocv.VideoWriter, scale float64) {
	ModifyVideo(image.ModifySaturation, videoIn, videoOut, scale)
}
*/
