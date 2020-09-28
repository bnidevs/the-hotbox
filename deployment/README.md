### frontend_clean.py

this is a post-deploy codepipeline stage lambda meant to move everything in the frontend folder into the root directory so the webpages display correctly

this function is allotted 10 seconds and 256mb of memory, if it takes longer/more memory due to the amount of files in the s3 bucket, switch to java
