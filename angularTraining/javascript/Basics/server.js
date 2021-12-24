const server=require('hhtp').createServer((req,response)=>{
    server.on('connection',(stream))=>{
        console.log("server connected")
    }
})