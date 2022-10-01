import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class DataService {
  public currUser: string = ''
  constructor() { }

  public saveCurrUser(user: string): void {
    this.currUser = user;
  }

  public sendUser(): string {
    return this.currUser;
  }
}
