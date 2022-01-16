import { Component, OnInit } from '@angular/core';
//  import { BallServiceService } from '../ball-service/ball-service.service';

export class ball{
  id?:number;
  description?:string;
  color?:string;
  clicked?:boolean;
}
@Component({
  selector: 'app-ball-display',
  templateUrl: './ball-display.component.html',
  styleUrls: ['./ball-display.component.css']
})
export class BallDisplayComponent implements OnInit{
 
  noOfBalls!:number;
  msg!:string;
  flag:boolean=true;
  balls:ball[]=[];
  Ans!:number;
  gameStatus!:boolean;
  noOfAttempts:number=0;
  attempts:number=0;
  remainingAttempts!:number;
  constructor() {
   
  }

  ngOnInit(): void {
    
  }
 
  getBall(){
    this.flag=false
    console.log("jiiii")
    console.log("no of balls",this.noOfBalls)
    for(let i=1;i<=this.noOfBalls;i++){ 
      this.balls.push({id:i,description:"ball"+i,color:"pink",clicked:false})
    }
    this.noOfAttempts=Math.round(Math.log2(this.noOfBalls))
    this.getRandNum()
  }
  getRandNum(){
    let a=Math.random()*this.noOfBalls
    this.Ans=Math.floor(a)+1;
    console.log("correct ans is ",this.Ans)
    return this.Ans
  }
  changeCS(givenball:ball){
    
    console.log("id",givenball.id)
    if(givenball.id! < this.Ans){

      givenball.color="green";
      this.gameStatus=false;
     
    }
    else if(givenball.id! > this.Ans){
      givenball.color="red";
      this.gameStatus=false;
      
    }
    else{
      givenball.color="blue";
      let msg="you won the game";
      alert(msg)
      this.msg=msg
      this.gameStatus=true;
     
    }
  }
  checkBall(givenball:ball){
  
    givenball.clicked=true;
    this.attempts=this.attempts+1
  
    console.log(givenball.clicked)
    this.remainingAttempts=this.noOfAttempts-this.attempts;
    if(this.remainingAttempts>=0){
      console.log("attempts in loop after click",this.remainingAttempts)
      this.changeCS(givenball);
    }else if(this.remainingAttempts<0){
      let msg="You lost all attempts"+" Correct number was "+this.Ans
      alert(msg)
      this.gameStatus=true
      this.msg=msg
      
      this.remainingAttempts=this.noOfAttempts
    }
    
  }
  
  restartGame(){
    
    this.balls=[];
    this.noOfAttempts=Math.round(Math.log2(this.noOfBalls));
    for(let i=1;i<=this.noOfBalls;i++){
      this.balls.push({id:i,description:"ball"+i,color:"pink"})
    }
    let a=Math.random()*this.noOfBalls
    this.Ans=Math.floor(a)+1;
    console.log("%c correct ans","background-color:pink",this.Ans);
    console.log(this.balls);
  }
  
  exit():void{
  
    this.balls=[]   
    this.noOfAttempts=0
    this.attempts=0
    this.Ans=0
    this.gameStatus=true
  }
  
}
  



  

  

