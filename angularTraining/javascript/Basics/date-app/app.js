
var bookArray=[];
function init(){
    if (localStorage.getItem("bookArray")){
        bookArray=JSON.parse(localStorage.getItem("bookArray"))
        for (var i=0;i<bookArray.length;i++){
            var bookName=bookArray[i]
        }
    }
}
function addTaskPressed(){
    var bookName=document.getElementById("bookname").value;
   // var bookObj ={bookname:bookName}
    bookArray.push(bookname);

}
localStorage.setItem("bookArray",JSON.stringify( bookArray));
// function onRegisterPressed(){
//     // var bookName=documenet.getElementById("bookname").value;
//     // var bookObj ={bookname:bookName}
//     // bookArray.push(bookname);

//     localStorage.booksRecord=JSON.stringify(bookArray);
//     var  table=document.getElementById("regtable")
//     var row=table.insertRow();
//     var bookCell=row.insertCell(0);
//     bookCell.innerHTML=bookName;
//     document.getElementById("bookName").value="";
//     localStorage.setItem(bookArray)

// }