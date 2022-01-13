import { Component, OnInit } from '@angular/core';
import { ProductService } from '../ProductService/product.service';
@Component({
  selector: 'app-prd-operation',
  templateUrl: './prd-operation.component.html',
  styleUrls: ['./prd-operation.component.css'],
  providers:  [ ProductService ]
})
export class PrdOperationComponent implements OnInit {
   Getproducts!:any
   newId:number=0
   newName:string=""
  //  prdAfterDel!:any
  //  prdAfterAdd!:any
   //prd!:any
  constructor(private productService:ProductService) {
    this.Getproducts=this.productService.GetProducts()
   }
  ngOnInit(): void {
    
    //this.productService.AddProduct(5,"teapack")
    //this.prdAfterAdd=this.productService.GetProducts()
    // this.productService.DeleteProduct(0)
    // this.prdAfterDel=this.productService.GetProducts()
   
  }
  deleteProduct(ID:any){
     this.productService.DeleteProduct(ID)
     this.Getproducts=this.productService.GetProducts()
  }
  
  addProduct(){
    var newProduct={id:this.newId,name:this.newName};
    this.productService.AddProduct(newProduct);
    console.log(newProduct)
  }

}