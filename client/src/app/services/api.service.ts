import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export type User = {
  [x:string]: any;
  username: string;
  email: string;
  password: string;
}
export type Song = {
  [x:string]: any;
  artist: string;
  title: string;
  album: string;
}

export type Playlist = {
  [x:string]: any;
  artist: string
  title: string
  album: string
  creator: string
  playlist_name: string
  songs: Song[]
}

export type LoginCreds = {
  username: string;
  password: string;
}
@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) { }

  public getCreds() {
    return this.http.get("localhost:5000/check-login")
  }

  public login(info: LoginCreds) {
    return this.http.post("localhost:5000/login", info)
  }

  public createUser(newUser: User) {
    return this.http.post("localhost:5000/new-user", newUser)
  }
}
