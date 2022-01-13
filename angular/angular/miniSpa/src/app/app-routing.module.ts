import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutComponent } from './about/about.component';
import { ContactComponent } from './contact/contact.component';
import { HomeComponent } from './home/home.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { XyzComponent } from './xyz/xyz.component';

const routes: Routes = 
[{ path: 'banking', loadChildren: () => import('./banking/banking.module').then(m => m.BankingModule) },
{path:'home',
children:[{path:'',component:HomeComponent,pathMatch:"full"},{path:'xyz',component:XyzComponent}]
},

{path:'',component:HomeComponent},
{path:'about',component:AboutComponent},
{path:'contact',component:ContactComponent},
{path:'**',component:NotFoundComponent}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
