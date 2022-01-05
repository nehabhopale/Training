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
     this.productService.DeleteProduct(0)
     this.Getproducts=this.productService.GetProducts()
  }
  
  addProduct(){
      this.productService.AddProduct(5,"teapack");
      this.Getproducts=this.productService.GetProducts()
  }

}