import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SellingComponent } from './selling/selling.component';
import { PrdOperationComponent } from './prd-operation/prd-operation.component';
//import { ProductService } from './ProductService/product.service';

@NgModule({
  declarations: [
    AppComponent,
    SellingComponent,
    PrdOperationComponent,
   // ProductService
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
   // ProductService
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
