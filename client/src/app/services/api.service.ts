import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export type User = {
  [x:string]: any;
  username: string;
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

  private playHost: string = "localhost:5000/playlist/"

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

  public addPlaylist(playlist: Playlist) {
    return this.http.post(this.playHost + 'create', playlist)
  }

  public getSinglePlaylist(playlistName: string) {
    return this.http.get("localhost:5000/find-playlist/" + playlistName )
  }

  public getAllUserPlaylists(username: string) {
    return this.http.get(this.playHost + "get-all/" + username);
  }

  public getAllSongs() {
    return this.http.get("localhost:5000/songs");
  }

  public addSong(song: Song) {
    return this.http.post(this.playHost + "add-song", song);
  }

  public deletePlaylist(playId: string) {
    return this.http.delete(this.playHost + "delete" + playId)
  }

  public deleteSong(id: string) {
    return this.http.delete(this.playHost + "delete-song" + id)
  }


}
