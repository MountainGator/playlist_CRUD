import { Component, OnInit } from '@angular/core';
import {MatDialog} from "@angular/material/dialog";
import {User} from "../../services/api.service";
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  public login: boolean = true;
  public userCreds: User = {
    username: '',
    email: '',
    password: ''
  }
  public create: boolean = false;
  public passwordMatch: string = '';
  constructor(private dialog: MatDialog, private api: ApiService) { }

  ngOnInit(): void {
  }

  public toggle(): void {
    this.create = !this.create;
  }

  public checkPwd(): void {

  }

  public createUser(): void {
    if(this.userCreds.password === this.passwordMatch) {
      this.api.createUser(this.userCreds).subscribe((res: any) => {

      })

    } else this.dialog.open()
  }

}
