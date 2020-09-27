var bucketName = "thehotboxupload";
var bucketRegion = 'us-east-1';
var IdentityPoolId = 'us-east-1:2271f583-09e5-4212-b72a-4024f2cea3c5';

AWS.config.update({
  region: bucketRegion,
  credentials: new AWS.CognitoIdentityCredentials({
    IdentityPoolId: IdentityPoolId
  })
});

var s3 = new AWS.S3({
  apiVersion: "2006-03-01",
  params: { Bucket: bucketName }
});

function uploadVideo() {
  var files = document.getElementById("upload").files;
  if (!files.length) {
    return;
  }
  var file = files[0];
  var fileName = file.name;

  var upload = new AWS.S3.ManagedUpload({
    params: {
      Bucket: bucketName,
      Key: fileName,
      Body: file
    }
  });

  var promise = upload.promise();

  promise.then();
}

document.getElementById("upload-btn").addEventListener("click", uploadVideo);
