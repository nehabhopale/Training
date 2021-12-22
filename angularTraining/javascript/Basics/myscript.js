//console.log("hello world")
document.getElementById("demo").innerHTML = "Hello World!";
const element = document.getElementById("id01");
element.innerHTML = "New Heading";

document.getElementById("id01").style.color="red";


///event to changeour inner html after click event
function changeText(id){
    id.innerHTML="thats new heading now";

}
document.getElementById("Btn").addEventListener("click", myFunction);

function myFunction() {
    alert ("Hello World!");
}

function NewData(id){
    const txt = '{"name":"neha", "age":30}'
    const obj = JSON.parse(txt);
    id.innerHTML = obj.name + ", " + obj.age;
}
const obj = {name: "neha", age: 40};
const myJSON = JSON.stringify(obj);
document.getElementById("newdemo").innerHTML = myJSON;