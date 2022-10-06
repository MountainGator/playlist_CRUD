import { Component, OnInit } from '@angular/core';
import {MatDialog} from "@angular/material/dialog";
import {User} from "../../services/api.service";
import { ApiService } from '../../services/api.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  public userCreds: User = {
    username: '',
    password: ''
  }
  public create: boolean = false;
  public passwordMatch: string = '';
  public wasError: boolean = false;
  constructor(private dialog: MatDialog, private api: ApiService, private router: Router) { }

  ngOnInit(): void {
  }

  public toggle(): void {
    this.create = !this.create;
  }

  public checkPwd(): void {
    this.api.login(this.userCreds).subscribe({
      next: res => {
        this.wasError = false;
        this.router.navigate([""]);
      },
      error: e => {
        console.error(e)
        this.wasError = true;
      }
    })
  }

  public createUser(): void {
    if(this.userCreds.password === this.passwordMatch) {
      this.api.createUser(this.userCreds).subscribe({
        next: (res: any) => {
          this.wasError = false;
          this.router.navigate([""])
        },
        error: e => {
          console.error(e)
          this.wasError = true;
        }
    })

    } 
  }

}
