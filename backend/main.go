package main

import (
	"fmt"
	"strings"
	"os"
	"gocv.io/x/gocv"
)

// method;
// adds/subtracts a constant value from each pixel, modifying the brightness
func ModifyBrightness(frame *gocv.Mat, change uint8, inc bool) {
	if change == 0 { return }

	// channels is of type []Mat, each of the three channels of frame
	// are now their own Mat type, and we can work with them separately
	channels := gocv.Split(*frame)

	// now we go through each channel and add the value of change to each pixel
	for i := 0; i < 3; i++ {
		// POSSIBLE OPTIMIZATION
		// Because we'll be doing multiple video operations at a time,
		// going through the frame as many times as we have adjustments to make is inefficient
		// if we could do this pixel-by-pixel we could do all of the adjustments at the same time
		// for now this is good but if we find our code to be too slow then we can come back to this

		if inc {
			channels[i].AddUChar(change) // channel i += change
		} else {
			channels[i].SubtractUChar(change) // channel i -= change
		}
	}

	// merges all the Mat's in channels into one multi-channel Mat, that being frame
	// now that we have added the value to the channels, we can put them back together in the frame
	gocv.Merge(channels, frame)

	// no return, the changes happen to the frame's reference
}

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
	defer out.Close()

	curr := gocv.NewMat() // reader mat
	defer curr.Close()
	
	for {
		
		// the frame curr is of type CV8UC3
		// CV8U means it stores unsigned chars (lit. 8 bit Unsigned)
		// C3 means it has 3 channels, these channels represent the Blue Green and Red respectively
		if ok := video.Read(&curr); !ok {
			fmt.Println("video reading stopped") // read frame to reader mat
			
			return
		}
		
		if curr.Empty() { 
			continue
		}
		
		// function call;
		// takes in a Mat*, a change value, and a boolean that asks if you're increasing or decreasing
		// changes the frame pointed to by the Mat*
		ModifyBrightness(&curr, 50, false) // this reduces curr's brightness by 50
		out.Write(curr)

	}

	// NOTE: the output doesn't have sound
}
