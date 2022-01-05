import { Component, OnInit } from '@angular/core';
import { ProductService } from '../ProductService/product.service';
@Component({
  selector: 'app-selling',
  templateUrl: './selling.component.html',
  styleUrls: ['./selling.component.css'],
  providers:  [ ProductService ]
})
export class SellingComponent implements OnInit {
  prd!:any
  constructor(private productService:ProductService) { }

  ngOnInit(): void {
    this.prd=this.productService.GetMaxSoldProduct()
  }
  
}
