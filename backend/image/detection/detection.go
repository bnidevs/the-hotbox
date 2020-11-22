package detection

import (
	"math"
	stdimage "image"
	"gocv.io/x/gocv"
)

/*
func DetectFace(img *gocv.Mat) []stdimage.Rectangle {
	harrcascade := "./haarcascade/haarcascade_frontalface_default.xml"
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(harrcascade)
	defer classifier.Close()

	return classifier.DetectMultiScale(*img)
}


// cut roughly 25% of height on the top of eyes
// func cutEyebrows()

func DetectEyes(img *gocv.Mat) []stdimage.Rectangle {
	harrcascade := "./haarcascade/haarcascade_eye.xml"
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(harrcascade)
	defer classifier.Close()

	eyeRects := classifier.DetectMultiScale(*img)
	eyesImg := 
	eyeRects := cutEyebrows(eyesImg)

	return eyeRects
}



// https://medium.com/@stepanfilonov/tracking-your-eyes-with-python-3952e66194a6#:~
// :text=OpenCV%20can%20put%20them%20in,left%20eye%20and%20vice%2Dversa.&text=The%20
// good%20thing%20about%20it,images(only%20two%20colors).

// https://towardsdatascience.com/real-time-eye-tracking-using-opencv-and-dlib-b504ca724ac6

// https://imotions.com/blog/eye-tracking-work/
// https://www.sr-research.com/about-eye-tracking/

func DetectPupils(img *gocv.Mat) []stdimage.Rectangle {
	eyeRects = DetectEyes(img)
}

*/