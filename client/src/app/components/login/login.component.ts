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
  public userClass = 'text-info';
  public passwordClass = 'text-info';
  public passwordText: string = 'Password';
  public confirmText: string = 'Confirm Password'
  public create: boolean = false;
  public passwordMatch: string = '';
  public wasError: boolean = false;
  public userLabel: string = 'Username';
  constructor(private dialog: MatDialog, private api: ApiService, private router: Router) { }

  ngOnInit(): void {
  }

  public toggle(): void {
    this.create = !this.create;
  }

  public checkPwd(): void {
    const userIsValid = this.validateUser();
    if(this.userCreds.password.length > 7 && userIsValid) {
      this.passwordClass = 'text-info';
      this.passwordText = "Password";
      this.api.login(this.userCreds).subscribe({
        next: res => {
          this.wasError = false;
          console.log("login response", res)
          this.router.navigate([""]);
        },
        error: e => {
          console.error(e)
          this.wasError = true;
        }
      })
    } else {
      this.passwordClass = 'text-danger';
      this.passwordText = "Password too short";
    }
  }

  public validateUser (): boolean{
    if(this.userCreds.username.length > 5) {
      this.userClass ="text-info";
      this.userLabel = 'Username';
      return true
    } else {
      this.userClass ="text-danger";
      this.userLabel = 'Username too short';
      return false
    }
  }
  public validatePwd(): boolean {
    if(this.userCreds.password.length > 7){
      this.passwordClass = 'text-info';
      this.passwordText = "Password";
      this.confirmText = "Confirm Password";
      if(this.userCreds.password === this.passwordMatch) {
        return true
      } else {
        this.passwordText = this.confirmText = "Passwords don't match";
        this.passwordClass = 'text-danger';
        return false
      }
    } else {
      this.passwordClass = 'text-danger';
      this.passwordText = this.confirmText = "Password too short";
      return false
    }
  }

  public createUser(): void {
    const pwdIsValid = this.validatePwd();
    const userIsValid = this.validateUser();
    if(pwdIsValid && userIsValid) {
        this.passwordText = 'Password';
        this.confirmText = "Confirm";
        this.passwordClass = 'text-info';
        this.userClass = "text-info";
        this.userLabel = 'Username';
        this.api.createUser(this.userCreds).subscribe({
          next: (res: any) => {
            console.log("create response", res)
            this.wasError = false;
            this.router.navigate([""])
          },
          error: e => {
            console.error(e)
            this.wasError = true;
          }
        })
    } else return
  }

}
