import { Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/services/api.service';
import { DataService } from 'src/app/services/data.service';
import { Subscription } from 'rxjs';
import { Playlist, Song } from 'src/app/services/api.service';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { MsgDialogComponent } from '../msg-dialog/msg-dialog.component';
import { ConfirmDialogComponent } from '../confirm-dialog/confirm-dialog.component';
import { AddSongDialogComponent } from '../add-song-dialog/add-song-dialog.component';
import { MatSnackBar } from '@angular/material/snack-bar';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnDestroy {
  private credSubscription!: Subscription;
  public currentUser: string = '';
  public newSong: Song = {
    artist: '',
    title: '',
    album: ''
  }
  public newPlaylist: Playlist = {
    creator: '',
    playlist_name: '',
    songs: []
  }
  public edit: boolean = false;
  public currPlayName: string = ''
  public playList: Playlist[] = [];
  public songsList: Song[] = [];

  constructor(
    private router: Router, 
    private api: ApiService, 
    private data: DataService, 
    private dialog: MatDialog,
    private snackBar: MatSnackBar
    ) { }

  ngOnInit(): void {
    this.getCreds();
    this.getSongs();
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
            this.getPlaylists();
          }
        },
        error: (e: Error) => {
          console.error(e)
          this.router.navigate(["/login"]);
        }
      }
    );
  }

  private getPlaylists() {
    this.api.getAllUserPlaylists(this.currentUser).subscribe({
      next: (res: {data: Playlist[]} | any) => {
        this.playList = res.data; 
      },
      error: e => {
        console.error("error finding user playlists", e)
      }
    })
  }

  private getSongs() {
    this.api.getAllSongs().subscribe({
      next: (res: {data: Song[]} | any) => {
        this.songsList = res.data;
      },
      error: e => {
        console.error("error finding songs", e)
      }
    })
  }

  public toggleEdit(name: string) {
    this.edit = !this.edit;
    this.currPlayName = this.edit === true ? name : '';
  }

  public createPlaylist() {
    this.newPlaylist.creator = this.currentUser;
    this.api.addPlaylist(this.newPlaylist).subscribe({
      next: () => {
        this.snack("Playlist created", "success-bar");
        this.getPlaylists();
      },
      error: e => {
        console.error("playlist create error", e)
        this.snack("Error creating Playlist", "danger-bar");
      }
    })
  }

  public addSubtractPlaylist(func: string, song: Song) {
    const ref: MatDialogRef<AddSongDialogComponent> = this.dialog.open(AddSongDialogComponent, {data: {func: func, song: song, list: this.playList}});

    ref.afterClosed().subscribe((res: {id: string, playlist: Playlist}) => {
        this.api.updatePlaylist(res).subscribe({
          next: () => {
            this.snack("Playlist updated!", "success-bar")
          },
          error: e => {
            console.error("update playlist error", e)
            this.snack("Error updating playlist", "danger-bar")
          }
        })
      }
    )
  }

  public addSong(){
    if(this.newSong.title === '' || this.newSong.artist === '' || this.newSong.album === '') {
      this.snack("Error. Missing information", "danger-bar")
      return
    } else {
      this.api.addSong(this.newSong).subscribe({
        next: (res: any) => {
        console.log('add song res:', res)
        this.snack("Song added!", "success-bar")
        },
        error: e => {
          console.error('add song error',e)
          this.snack("Error adding song", "danger-bar")
        }
      })
    }
  }

  public deleteSong(id: string) {
    const ref: MatDialogRef<ConfirmDialogComponent> = this.dialog.open(ConfirmDialogComponent);
    ref.afterClosed().subscribe(res => {
      if (res) {
        this.api.deleteSong(id).subscribe()
        this.snack("Song deleted!", "success-bar")
      } else return
    })
  }

  private snack(msg: string, color: string): void {
    this.snackBar.open(msg, '', {horizontalPosition: 'end', verticalPosition: 'top', duration: 2000, panelClass: [color]})
  }

}
