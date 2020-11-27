package detection

import (
	stdimage "image"
	"gocv.io/x/gocv"
	// "fmt"
)


func DetectFaces(img *gocv.Mat, faceClassifier *gocv.CascadeClassifier) []stdimage.Rectangle {
	return faceClassifier.DetectMultiScale(*img)
}


func DetectEyes(img *gocv.Mat, eyeClassifier *gocv.CascadeClassifier, 
				faceClassifier *gocv.CascadeClassifier) []stdimage.Rectangle {
	faces := faceClassifier.DetectMultiScale(*img)
	eyes := eyeClassifier.DetectMultiScale(*img)

	// Remove false positives
	if len(faces) == 0 {
		eyes = make([]stdimage.Rectangle,0)
		return eyes
	}

	// We only need the upper half of a face since that's where eyes are
	for i:=0; i<len(faces); i++ {
		faces[i].Max.Y = (faces[i].Max.Y + faces[i].Min.Y) / 2
	}
	
	for i:=0; i<len(eyes); i++ {
		isInFace := false
		for j:=0; j<len(faces); j++ {
			if eyes[i].In(faces[j]) {
				isInFace = true
				break
			}
		}
		if (!isInFace) {
			eyes[i], eyes[len(eyes)-1] = eyes[len(eyes)-1], eyes[i]
			eyes = eyes[:len(eyes)-1]
			i--
		}
	}

	return eyes
}



// https://medium.com/@stepanfilonov/tracking-your-eyes-with-python-3952e66194a6#:~
// :text=OpenCV%20can%20put%20them%20in,left%20eye%20and%20vice%2Dversa.&text=The%20
// good%20thing%20about%20it,images(only%20two%20colors).

// https://towardsdatascience.com/real-time-eye-tracking-using-opencv-and-dlib-b504ca724ac6

// https://imotions.com/blog/eye-tracking-work/
// https://www.sr-research.com/about-eye-tracking/

// get the center(KeyPoint.X and .Y) and diameter(KeyPoint.Size) of pupils, 
// and the center point of each eye
func DetectPupils(img *gocv.Mat, eyeClassifier *gocv.CascadeClassifier, 
				faceClassifier *gocv.CascadeClassifier) ([]gocv.KeyPoint, []stdimage.Point) {
	eyes := DetectEyes(img, eyeClassifier, faceClassifier)

	eyeCenters := make([]stdimage.Point,len(eyes))

	for i:=0; i<len(eyes); i++ {
		eyeCenters[i].X = eyes[i].Min.X + (eyes[i].Max.X-eyes[i].Min.X)/2
		eyeCenters[i].Y = eyes[i].Min.Y + (eyes[i].Max.Y-eyes[i].Min.Y)/2

		// Cut roughly 25% of height on the top of eyes
		eyes[i].Min.Y = eyes[i].Min.Y+(eyes[i].Max.Y-eyes[i].Min.Y)/4

	}

	detectorParams := gocv.NewSimpleBlobDetectorParams()
	detectorParams.SetFilterByArea(true)
	detectorParams.SetMaxArea(1500.0)
	pupilDetector := gocv.NewSimpleBlobDetectorWithParams(detectorParams)


	eyeKeyPts := make([]gocv.KeyPoint,len(eyes))
	imgData := img.DataPtrUint8()
	// rows := img.Rows()
	cols := img.Cols()
	for i:=0; i<len(eyes); i++ {
		eyeCol := eyes[i].Max.X - eyes[i].Min.X
		eyeRow := eyes[i].Max.Y - eyes[i].Min.Y
		baseCol := eyes[i].Min.X
		baseRow := eyes[i].Min.Y

		// extract the eyes from image
		eyeImg := gocv.NewMatWithSize(eyeRow, eyeCol, gocv.MatTypeCV8UC1)
		eyeData := eyeImg.DataPtrUint8()
		for j:=0; j<eyeRow; j++ {
			for k:= 0; k<eyeCol; k++ {
				eyeData[j*eyeCol+k] = imgData[(baseRow+j)*cols+(baseCol+k)]
			}
		}

		// Process the gray eye image so that the blob is easier to detect
		gocv.Threshold(eyeImg, &eyeImg, 52, 255,gocv.ThresholdBinary)
		tmpMat := gocv.NewMat()
		for j:=0; j<2; j++ {
			gocv.Erode(eyeImg, &eyeImg, tmpMat)
		}
		for j:=0; j<4; j++ {
			gocv.Dilate(eyeImg, &eyeImg, tmpMat)
		}
		gocv.MedianBlur(eyeImg, &eyeImg, 5)
		
		// display modified image for testing
		// if i==0 {
		// 	win := gocv.NewWindow("hello")
		// 	win.IMShow(eyeImg)
		// 	win.WaitKey(1000)
		// }
		
		keyPts := pupilDetector.Detect(eyeImg)
		// Calculate the position of pupils on the original image
		if len(keyPts)>0 {
			eyeKeyPts[i] = keyPts[0]
			eyeKeyPts[i].X += float64(baseCol)
			eyeKeyPts[i].Y += float64(baseRow)
			// eyeKeyPts[i].Size *= 5
		}
	}

	return eyeKeyPts, eyeCenters
}

