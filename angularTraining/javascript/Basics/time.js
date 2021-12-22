// The setTimeout() is executed only once.

// If you need repeated executions, use setInterval() instead.
const timeOut=setTimeout(myName,5000) //call after 5 sec
clearTimeout(timeOut) //prevents myname from starting
function myName(){
    console.log("neha")
}
//Use the clearTimeout() method to prevent the function from starting.

interval=setInterval(myName, 1000);//it will call my name every sec
clearInterval(interval)