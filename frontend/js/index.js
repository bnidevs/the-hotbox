var bucketName = "thehotboxupload";
var bucketRegion = 'us-east-1';
var IdentityPoolId = 'us-east-1:2271f583-09e5-4212-b72a-4024f2cea3c5';

let uploadprogbar = document.getElementById("uploadprogressbar");
let uploadprognum = document.getElementById("uploadprogress");

var videofilename;

AWS.config.update({
  region: bucketRegion,
  credentials: new AWS.CognitoIdentityCredentials({
    IdentityPoolId: IdentityPoolId
  })
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
    uploadprogbar.value = e.loaded / filesize * 100;
    uploadprognum.innerText = shorten_float(uploadprogbar.value) + "%";
  });

  var promise = upload.promise();

  promise.then(
    function(data) {
      console.log("success");
      document.getElementById("upload-label").innerText = "Upload Done";
      checkFileSize();
      callLambdaProcess();
    },
    function(err) {console.log("error")}
  );
}

function checkFileSize() {
  var paramspayload = {"videofilename":videofilename}

  var lambdaParams = {
    FunctionName: '035225278288:function:thehotboxcheckfilesize',
    Payload: JSON.stringify(paramspayload)
  };
  var lambda = new AWS.Lambda({apiVersion: '2015-03-31'});
  lambda.invoke(lambdaParams, function(err, data){
    if(err) console.log(err, err.stack);
    else console.log(data);
  });
}

function callLambdaProcess() {
  var paramspayload = {"videofilename":videofilename}

  var lambdaParams = {
    FunctionName: '035225278288:function:thehotboxvideoprocess',
    Payload: JSON.stringify(paramspayload)
  };
  var lambda = new AWS.Lambda({apiVersion: '2015-03-31'});
  lambda.invoke(lambdaParams, function(err, data){
    if(err) console.log(err, err.stack);
    else console.log(data);
  });
}

document.getElementById("upload-btn").addEventListener("click", uploadVideo);

var shorten_float = (f) => {
  return f.toFixed(2);
}

function displaySliderValue(id1, id2){
  var slider_value = document.getElementById(id1);
  var slider_dragger = document.getElementById(id2);
  slider_value.innerHTML = slider_dragger.value + " ";
}

function toggleTheme(){
  var element = document.body;
  element.classList.toggle("dark-mode");
}
