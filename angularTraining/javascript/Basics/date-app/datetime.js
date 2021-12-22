const myInterval = setInterval(myTimer, 1000);

function myTimer() {
  const date = new Date();
  document.getElementById("demo").innerHTML = date.toLocaleTimeString();
  document.getElementById("newdemo").innerHTML = date.toLocaleDateString();
}
