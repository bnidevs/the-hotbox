package video

import (
	//"fmt"
	"gocv.io/x/gocv"
	"../image"
)

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
	curr := gocv.NewMat() // reader mat
	defer curr.Close()

	for {
		if ok := videoIn.Read(&curr); !ok {
			return
		}
		
		if curr.Empty() { 
			continue
		}
		
		image.ModifyContrast(&curr, alpha)
		videoOut.Write(curr)

	}
}
	

func ModifySaturation(videoIn *gocv.VideoCapture, videoOut *gocv.VideoWriter, scale float64) {
	curr := gocv.NewMat() // reader mat
	defer curr.Close()

	for {
		if ok := videoIn.Read(&curr); !ok {
			return
		}
		
		if curr.Empty() { 
			continue
		}
		
		image.ModifySaturation(&curr, scale)
		videoOut.Write(curr)

	}
}
