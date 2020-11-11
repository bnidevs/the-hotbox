import numpy as np

class Parameters:
	def __init__(self, brightness=0, contrast=0, saturation=0, \
				distortion=0, noise=0, currFrame=0):
		self.brightness = brightness
		self.contrast = contrast
		self.saturation = saturation
		self.distortion = distortion
		self.noise = noise
		self.currFrame = currFrame

# nice closure to expedite the process of keeping the values between 0 and 255
def toUint8(val):
	if val < 0:
		return 0
	elif val > 255:
		return 255
	else:
		return np.uint8(val)