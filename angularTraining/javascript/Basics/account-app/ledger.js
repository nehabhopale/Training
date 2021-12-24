const fs=require('fs')
class ledger{
    constructor(Account){
        this.account=Account
    }
    listenCall(){
        this.account.on('amount-deposited',function(a,b,c){
            console.log("Event caught",a,b,c);
            //writefile
            var d={}
            console.log("Event caught",a,b,c);
            d["balance"]=a
            d["accNo"]=b
            d["transaction"]=c
            fs.writeFile('./file.txt',JSON.stringify(d)+'\n',{
                encoding: "utf8",
                flag: "a",
                mode: 0o666
                },(err)=>{
                if(err){
                    console.log(err);
                }
            });
        })    
        this.account.on('amount-withdrawan',function(a,b,c){
            var w={}
            console.log("Event caught",a,b,c);
            w["balance"]=a
            w["accNo"]=b
            w["transaction"]=c

           // console.log(JSON.stringify(o))
            fs.writeFile('./file.txt',JSON.stringify(w)+'\n',{
                encoding: "utf8",
                flag: "a",
                mode: 0o666
                },(err)=>{
                if(err){
                    console.log(err);
                }
            });
        })     
    }
    // this.account.on('amount-deposited',(a,b,c)=> {

        //     console.log(a, b,c);
        // });

        // this.account.on('amount-withdrawan',(a,b,c)=> {

        //     console.log(a, b,c);
        // });


}
module.exports=ledger