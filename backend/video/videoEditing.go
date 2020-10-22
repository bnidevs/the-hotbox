package video

import (
	//"fmt"
	"container/heap"
	"sync"

	"../image"
	"../utils"
	"gocv.io/x/gocv"
)

// ModifyVideo helps refactor the code by extracting the goroutine implementation
func ModifyVideo(imagefunc func(*gocv.Mat, float64), videoIn *gocv.VideoCapture, videoOut *gocv.VideoWriter, num float64) {
	curr := gocv.NewMat() // reader mat
	defer curr.Close()

	var i int = 0
	var wg sync.WaitGroup
	pq := make(utils.PriorityQueue, 1)
	heap.Init(&pq)

	// Create a goroutine to modify each frame simultaneously
	for ok := videoIn.Read(&curr); !ok; ok, i = videoIn.Read(&curr), i+1 {
		if curr.Empty() {
			continue
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			imagefunc(&curr, num)
			frame := &utils.Frame{
				Image:    curr,
				Priority: i,
			}
			heap.Push(&pq, frame)
		}(i)
	}
	wg.Wait()

	// Write the modified frames into the output video in the same order as the original
	for pq.Len() > 0 {
		frame := heap.Pop(&pq).(*utils.Frame)
		videoOut.Write(frame.Image)
	}
}

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
