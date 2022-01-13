import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SellingComponent } from './selling/selling.component';
import { PrdOperationComponent } from './prd-operation/prd-operation.component';

import { FormsModule } from '@angular/forms';
import { CheckerDirective } from './directives/checker.directive';
//import { ProductService } from './ProductService/product.service';

@NgModule({
  declarations: [
    AppComponent,
    SellingComponent,
    PrdOperationComponent,
    CheckerDirective,
   // ProductService
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule
   // ProductService
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
