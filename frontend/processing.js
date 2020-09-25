let progbar = document.getElementById("file");
let prognum = document.getElementById("progress");

async function progress() {
  while(progbar.value < 100){
    await new Promise(resolve => setTimeout(resolve, 100));
    progbar.value++;
    prognum.innerText = progbar.value + "%";
  }
}

window.onload = () => {
  progress();
}
