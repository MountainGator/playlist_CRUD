import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/services/api.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  private credSubscription!: Subscription;
  public currentUser: string = '';
  constructor(private router: Router, private api: ApiService) { }

  ngOnInit(): void {

  }
  public getCreds() {
    let tempUser: string | null = localStorage.getItem("current user")

    if(!tempUser) {
      this.credSubscription = this.api.getCreds().subscribe(
        {
          next: (res: any) => {
            console.log('check login res', res)
            if (res.error){
              this.router.navigate(["/login"]);
            } else if(res.user) {
              localStorage.setItem('current user',JSON.stringify(res.user));
              this.currentUser = res.user;
            }
          },
          error: (e: Error) => {
            console.error(e)
            this.router.navigate(["/login"]);
          }
        }
      );
    } else {
      this.currentUser = tempUser;
    }

  }
}
