import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { WelcomeAppRoutingModule } from './welcomeapp-routing.module';
import { WelcomeAppComponent } from './welcomeapp.component';
import { TestComponent } from './my-test/mytest.component';

@NgModule({
  declarations: [
    WelcomeAppComponent,
    TestComponent
  ],
  imports: [
    BrowserModule,
    WelcomeAppRoutingModule
  ],
  providers: [],
  bootstrap: [WelcomeAppComponent]
})
export class WelcomeAppModule { }
