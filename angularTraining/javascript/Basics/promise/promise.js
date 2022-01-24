let userData=[
    {
    "name":"neha",
    "age":22
    },
    {
    "name":"pooja",
    "age":21
    },
    {
    "name":"sweety",
    "age":22
    },
]
let hobbyData=[
    {
    "name":"neha",
    "hobby":"running"
    },
    // {
    // "name":"pooja",
    // "hobby":"sleeping"
    // },
    {
    "name":"neha",
    "hobby":"dancing"
    },
]
// function onLoad(data,succeessCallBack,errCallBack){
//     if (data.length==3){
//         succeessCallBack(data)
//     }else{
//         errCallBack("couldn't load user data")
//     }
// }
// onLoad(userData,(obj)=>{
//     console.log("in success ",obj)
//     onLoad(hobbyData,(obj)=>{
//         console.log("in success ",obj)
//     },(err)=>{
//         console.log("in err",err)
//     })
// },(err)=>{
//     console.log("in err",err)
// })
function onUserLoadPromise(userdata){
    return new Promise((resolve,reject)=>{
        if(userdata.length==3){
            resolve(userdata)
        }else{
            reject("could't load user data")
        }
    })
}
function onhobbyLoadPromise(hobbydata){
    return new Promise((resolve,reject)=>{
        if(hobbydata.length==3){
            resolve(hobbydata)
        }else{
            reject("could't load hobby data")
        }
    })
}
const load=async()=>{
    var user =await onUserLoadPromise(userData)
    var hobby =await onhobbyLoadPromise(hobbyData)
    return [user,hobby]
}
load().then((obj)=>{
       console.log("in success",obj)
    })
    .catch((err)=>{
        console.log("in err",err)
    })

    

// onLoadPromise(userData)
// .then((obj)=>{
//     console.log("in success",obj)
//     return onLoadPromise(hobbyData)
// })
// .then((obj)=>{
//     console.log("inner succ",obj)
// })