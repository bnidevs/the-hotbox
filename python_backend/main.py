import cv2 
import sys

from video import videoEditing, videoIO


if __name__ == '__main__':
    if len(sys.argv) != 2:
        print(f"Argument error\nUsage: {sys.argv[0]} filepath")
        sys.exit(1)

    videoIn = cv2.VideoCapture(sys.argv[1])
    
    if videoIn.isOpened() is False:
        print(f"Error\nUnable to read video")

    


    # CLOSE EVERYTHING
    videoIn.release()
