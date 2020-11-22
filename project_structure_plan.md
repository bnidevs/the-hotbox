# Project Structure

***

- the-hot-box  
  - documentation and license
  - backend  
    - main.go
    - image
      - imageEditing.go
      - detection
        - detection.go
        - haarcascades (storing xml files for detection)
    - video
      - videoEditing.go
      - videoIO.go
    - utils
      - perlin
        - perlin.go
      - common.go
      - priorityqueue.go
    - api.go (probably need this one)  

***

## main.go

main.go should be simply calling apis or fucntions from other modules.  

## imageEditing.go

Frame editing is basiclly image editing. Functions in this module should only deal with one frame or image.  

## detection.go

Implement eye, pupil and face detection here.  

## videoEditing.go

We can implement all video effects with asynchronous or synchronous frame editing here.  

## videoIO.go

videoIO.go deals with reading and writing video files.  

## common.go

Commonly used functions can be put here. For example, the sub-function in ModifyBrightness3 can be put here.  

## perlin.go

Generate 1D, 2D or 3D Perlin noise.  

## priorityqueue.go

Implement priority queue here.  
