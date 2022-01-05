import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { WelcomeAppRoutingModule } from './welcomeapp-routing.module';
import { WelcomeAppComponent } from './welcomeapp.component';
import { TestComponent } from './my-test/mytest.component';
import { FooterComponent } from './footer/footer.component';
import { HeaderComponent } from './header/header.component';
import { SpinnerComponent } from './spinner/spinner.component';
import { ChangeimageComponent } from './changeimage/changeimage.component';
import { LoopingComponent } from './looping/looping.component';
import { GreetingComponent } from './greeting/greeting.component';
import { TwoWayComponent } from './two-way/two-way.component';
import { FormsModule } from '@angular/forms';
import { RadioComponent } from './radio/radio.component';
import { OutputFormatComponent } from './output-format/output-format.component';
import { StudentsComponent } from './students/students.component';
import { TestDirective } from './directives/test.directive';
import{TestingDirective}from './directives/testing.directive'


@NgModule({
  declarations: [
    WelcomeAppComponent,
    TestComponent,
    FooterComponent,
    HeaderComponent,
    SpinnerComponent,
    ChangeimageComponent,
    LoopingComponent,
    GreetingComponent,
    TwoWayComponent,
    RadioComponent,
    OutputFormatComponent,
    StudentsComponent,
    TestDirective,
    TestingDirective
  ],
  imports: [
    BrowserModule,FormsModule,
    WelcomeAppRoutingModule
  ],
  providers: [],
  bootstrap: [WelcomeAppComponent]
})
export class WelcomeAppModule { }
