import { Injectable } from '@angular/core';
export interface Products {
  id:number,
  name:string,
}

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  product!:Products[]

  constructor() {
    this.product = [
      { id:1,
        name:"mobile"
      },
      {id:2,
        name:"tv"
       
      },
      { id:3,
        name:"fridge"
      },
      {id:4,
        name:"shirt"
       
      },
    ]
   }
  GetProducts(){
     return this.product
  }
  AddProduct(Id:number,Name:string){
   var newPrd={id:Id,name:Name}
   this.product.push(newPrd)
  }
  GetMaxSoldProduct(){
    var number=this.getRandomInt(1,4)
    console.log(number)
    return this.product[number]
  }

  DeleteProduct(pos:number): void{
    this.product.splice(pos,1)
  }

  getRandomInt(min:number, max:number) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}
}
