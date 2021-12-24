const EventEmmiter=require('events')

const someChange=require('./myEmmitter')

const emmitter=new EventEmmiter()

emmitter.on("someEvent",function(){
    console.log("event caught")
})
someChange()
