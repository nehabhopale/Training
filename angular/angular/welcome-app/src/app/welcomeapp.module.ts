import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { WelcomeAppRoutingModule } from './welcomeapp-routing.module';
import { WelcomeAppComponent } from './welcomeapp.component';
import { TestComponent } from './my-test/mytest.component';
import { FooterComponent } from './footer/footer.component';
import { HeaderComponent } from './header/header.component';
import { SpinnerComponent } from './spinner/spinner.component';

@NgModule({
  declarations: [
    WelcomeAppComponent,
    TestComponent,
    FooterComponent,
    HeaderComponent,
    SpinnerComponent
  ],
  imports: [
    BrowserModule,
    WelcomeAppRoutingModule
  ],
  providers: [],
  bootstrap: [WelcomeAppComponent]
})
export class WelcomeAppModule { }
