package utils

type Parameters struct {
	Brightness int16
	Contrast float64
	Saturation float64
}

// nice closure to expedite the process of keeping the values between 0 and 255
func Int16ToUint8(val int16) uint8 { 
	if val < 0 {
		return 0
	} else if val > 255 {
		return 255
	} else {
		return uint8(val)
	}
}

func Float64ToUint8(val float64) uint8 { 
	if val < 0 {
		return 0
	} else if val > 255 {
		return 255
	} else {
		return uint8(val)
	}
}

// nice closure that returns the number and color that is the max of the three
func Max(b, g, r uint8) uint8 {
	if b > g && g > r {
		return b
	} else if g > r {
		return g
	} else {
		return r
	}
}
