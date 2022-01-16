import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { SliderModule } from '@syncfusion/ej2-angular-inputs';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FormsModule } from '@angular/forms';
import { BallDisplayComponent } from './ball-display/ball-display.component';
import {MatSliderModule} from '@angular/material/slider';
import { MatCheckboxModule} from '@angular/material/checkbox';


@NgModule({
  declarations: [
    AppComponent,
    BallDisplayComponent
  ],
  imports: [
    BrowserModule,FormsModule,MatSliderModule,
    SliderModule,BrowserAnimationsModule,MatCheckboxModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
