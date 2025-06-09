import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {LoginComponent} from "./components/login/login.component";
import {HomeComponent} from "./components/home/home.component";
import {CalenderComponent} from "./components/calender/calender.component";
import {TodoComponent} from "./components/todo/todo.component";
import {UploadComponent} from "./components/upload/upload.component";

const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full'
  },
  {
    path: 'login',
    component:LoginComponent
  },
  {
    path: 'home',
    component:HomeComponent
  }
  ,
  {
    path: 'calender',
    component:CalenderComponent
  }
  ,
  {
    path: 'todo',
    component:TodoComponent
  }
  ,
  {
    path: 'upload',
    component:UploadComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
