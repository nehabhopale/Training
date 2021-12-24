interface Person {
    firstName: string;
    lastName: string;
    age ?:number;
}
   
class UserP{
    firstName: string;
    lastName: string; 
    age ?:number;
}
var personList :Person[]=[];
var user11=new UserP()
user11.firstName="neha"
user11.lastName="b"

var user12=new UserP()
user12.firstName="pooja"
user12.lastName="c"
user12.age=12

personList.push(user11)
personList.push(user12)

function printDetails(person:Person[]) {
    for(let p of person){
        console.log(`firstname is ${p.firstName}`)
        console.log(`lastname is ${p.lastName}`)
        console.log(`age is ${p.age}`)
        console.log(`p is ${p}`) //shows [object object] with string interpolation
        console.log(p) 

        
    }
    console.log("*****USE OF FOR IN****")
    for (let q in person){
        console.log(`q is ${q}`)    //gives 0 and 1 that is indexes 
    }
    console.log("----anothere for way----- ")
   
    for(let k in person){
        console.log(person[k])
        console.log(`firstName of person is ${person[k].firstName}`)
        console.log(`lastName of person is ${person[k].lastName}`)
    }
}
printDetails(personList);