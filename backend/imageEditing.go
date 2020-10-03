package imageEditing

import (
	"fmt"
	"gocv.io/x/gocv"
)


// adds/subtracts a constant value from each pixel, modifying the brightness
func ModifyBrightness1(frame *gocv.Mat, change uint8, inc bool) {
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


func ModifyBrightness(frame *gocv.Mat, change int16) {
	framedata := frame.DataPtrUint8()

	// nice closure to expedite the process of keeping the values between 0 and 255
	limit := func(val int16) uint8 { 
		if val < 0 {
			return 0
		} else if val > 255 {
			return 255
		} else {
			return uint8(val)
		}
	}

	// framedata is one long array of uint8's, every third item represents a new pixel
	// and the three in between are the BGR channels
	for i := 0; i < len(framedata); i += 3 {
		// done like this so we can add weights
		framedata[i] = limit(int16(framedata[i]) + change) // B 
		framedata[i+1] = limit(int16(framedata[i+1]) + change) // G
		framedata[i+2] = limit(int16(framedata[i+2]) + change) // R
	}
}