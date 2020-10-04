# Project Structure

***

- the-hot-box  
  - documentation and license
  - backend  
    - main.go
    - imageEditing.go
    - videoEditing.go
    - videoIO.go
    - common.go
    - api.go (probably need this one)
    - more potential modules ...  

***

## main.go

main.go should be simply calling apis or fucntions from other modules.  

## imageEditing.go

Frame editing is basiclly image editing. Functions in this module should only deal with one frame or image.  

## videoEditing.go

We can implement all video effects with asynchronous or synchronous frame editing here.  

## videoIO.go

videoIO.go deals with reading and writing video files.  

## common.go

commonly used functions can be put here. For example, the sub-function in ModifyBrightness3 can be put here.  
