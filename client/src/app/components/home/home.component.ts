import { Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/services/api.service';
import { DataService } from 'src/app/services/data.service';
import { Subscription } from 'rxjs';
import { Playlist, Song } from 'src/app/services/api.service';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnDestroy {
  private credSubscription!: Subscription;
  public currentUser: string = '';

  public playList: Playlist[] = [];
  public songsList: Song[] = [];

  constructor(private router: Router, private api: ApiService, private data: DataService) { }

  ngOnInit(): void {
    this.getCreds();
  }

  ngOnDestroy(): void {
    this.credSubscription.unsubscribe();
  }

  public getCreds(): void {
      this.credSubscription = this.api.getCreds().subscribe(
        {
          next: (res: any) => {
            console.log('check login res', res)
            if (res.error){
              this.router.navigate(["/login"]);
            } else if(res.user) {
              this.currentUser = res.user;
            }
          },
          error: (e: Error) => {
            console.error(e)
            this.router.navigate(["/login"]);
          }
        }
      );

  }
}
