var bucketName = "thehotboxupload";
var bucketRegion = 'us-east-1';
var IdentityPoolId = 'us-east-1:2271f583-09e5-4212-b72a-4024f2cea3c5';

let progbar = document.getElementById("progbar");
let prognum = document.getElementById("progress");

var videofilename;

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

function uuidv4() {
  return ([1e7]+-1e3+-4e3+-8e3+-1e11).replace(/[018]/g, c =>
    (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
  );
}

function uploadVideo() {
  var files = document.getElementById("upload").files;
  if (!files.length) {
    return; // error message no file
  }

  var filesize = 0;
  for (var i = 0; i < files.length; i++){
    filesize += files[i].size;
  }

  if (filesize > 250000000){
    return; // error message file too big
  }

  var file = files[0];
  videofilename = uuidv4() + file.name;

  document.getElementById("progress-container").style.display = "flex";

  var upload = new AWS.S3.ManagedUpload({
    params: {
      Bucket: bucketName,
      Key: videofilename,
      Body: file
    }
  }).on("httpUploadProgress", function(e){
    progbar.value = e.loaded / filesize * 100;
    prognum.innerText = progbar.value + "%";
  });

  var promise = upload.promise();

  promise.then(
    function(data) {
      console.log("success");
      document.getElementById("upload-label").innerText = "Upload Done";
      callLambdaProcess();
    },
    function(err) {console.log("error")}
  );
}

function callLambdaProcess() {
  var lambdaParams = {
    FunctionName: '035225278288:function:thehotboxvideoprocess',
    Payload: '{"videofilename":"' + videofilename + '"}'
  };
  var lambda = new AWS.Lambda({apiVersion: '2015-03-31'});
  lambda.invoke(lambdaParams, function(err, data){
    if(err) console.log(err, err.stack);
    else console.log(data);
  });
}

document.getElementById("upload-btn").addEventListener("click", uploadVideo);



// Function shortens f to 2 decimal places
var shorten_float = (f) => {
  var shortened_string = f.toFixed(2);
  return shortened_string;
}
