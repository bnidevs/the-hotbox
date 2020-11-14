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
      getImgPreview();
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

function getImgPreview() {
  var paramspayload = {"videofilename":videofilename}

  document.getElementById("preview-image").src="assets/no_preview.png";

  var lambdaParams = {
    FunctionName: '035225278288:function:thehotboximagepreview',
    Payload: JSON.stringify(paramspayload)
  };
  var lambda = new AWS.Lambda({apiVersion: '2015-03-31'});
  lambda.invoke(lambdaParams, function(err, data){
    if(err) console.log(err, err.stack);
    else console.log(data);
  });
}

function callLambdaProcess() {
  var paramspayload = {
    "videofilename"     :   videofilename,
    "brightness_val"    :   document.getElementById("brightness-slider").value,
    "saturation_val"    :   document.getElementById("saturation-slider").value,
    "contrast_val"      :   document.getElementById("contrast-slider").value,
    "noise_val"         :   document.getElementById("noise-slider").value,
    "laser_eyes_check"  :   document.getElementById("laser-eyes").checked,
    "head_bulge_check"  :   document.getElementById("head-bulge").checked
  }

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

function displaySliderValue(id1, id2){
  var slider_value = document.getElementById(id1);
  var slider_dragger = document.getElementById(id2);
  slider_value.innerHTML = slider_dragger.value + " ";
}

function setDefaults(){
  document.getElementById("brightness-value").innerHTML = "50 ";
  document.getElementById("saturation-value").innerHTML = "50 ";
  document.getElementById("contrast-value").innerHTML = "50 ";
  document.getElementById("noise-value").innerHTML = "0 ";
  document.getElementById("brightness-slider").value = 50;
  document.getElementById("saturation-slider").value = 50;
  document.getElementById("contrast-slider").value = 50;
  document.getElementById("noise-slider").value = 0;
  document.getElementById("laser-eyes").checked = false;
  document.getElementById("head-bulge").checked = false;
}

function toggleTheme(){
  var element = document.body;
  var themeID = document.getElementById("themeSwitch").value;
  if(themeID == 1){
    document.documentElement.className="light-mode";
    document.getElementsByTagName("h1")[0].style.color = "black";
    document.getElementsByTagName("footer")[0].style.color = "black";
  }
  else if(themeID == 2){
    document.documentElement.className="dark-mode";
    document.getElementsByTagName("h1")[0].style.color = "white";
    document.getElementsByTagName("footer")[0].style.color = "white";
  }
  else if(themeID == 3){
    document.documentElement.className="halloween";
    document.getElementsByTagName("h1")[0].style.color = "#FF7000";
    document.getElementsByTagName("footer")[0].style.color = "#FF7000";
  }
  else if(themeID == 4){
    document.documentElement.className="thanksgiving";
    document.getElementsByTagName("h1")[0].style.color = "#FAF16A";
    document.getElementsByTagName("footer")[0].style.color = "FAF16A";
  }
  else if(themeID == 5){
    document.documentElement.className="christmas";
    document.getElementsByTagName("h1")[0].style.color = "#00cc00";
    document.getElementsByTagName("footer")[0].style.color = "white";
  }
  else if(themeID == 6){
    document.documentElement.className="newyears";
    document.getElementsByTagName("h1")[0].style.color = "hotpink";
    document.getElementsByTagName("footer")[0].style.color = "white";
  }
}

function checkDate(){
  var date = new Date();
  var month = date.getMonth()+1;
  if(month == 10){
    document.getElementById("themeSwitch").value=3;
  }
  else if(month == 11){
    document.getElementById("themeSwitch").value=4;
  }
  else if(month == 12){
    document.getElementById("themeSwitch").value=5;
  }
  else if(month == 1){
    document.getElementById("themeSwitch").value=6;
  }
  else{
    return;
  }
  toggleTheme();
}

var shorten_float = (f) => {
  return f.toFixed(2);
}

// Event Listeners
document.getElementById("upload-btn").addEventListener("click", uploadVideo);
window.addEventListener("beforeunload", setDefaults());
