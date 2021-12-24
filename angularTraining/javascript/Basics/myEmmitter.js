const EventEmmiter=require('events')


class MyEmmitter extends EventEmmiter{
    someChange(message){
        this.emit("someupdate","hi",123)
        console
    }
}
module.exports=MyEmmitter