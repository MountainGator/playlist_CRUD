import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Playlist, Song } from 'src/app/services/api.service';

interface DialogData {
  func: string;
  song: Song, 
  list: Playlist[]
}

@Component({
  selector: 'app-add-song-dialog',
  templateUrl: './add-song-dialog.component.html',
  styleUrls: ['./add-song-dialog.component.scss']
})
export class AddSongDialogComponent implements OnInit {
  private newList!: Playlist;

  constructor(public dialogRef: MatDialogRef<AddSongDialogComponent>, @Inject(MAT_DIALOG_DATA) public data: DialogData) { }

  ngOnInit(): void {
  }

  public updatePlaylist(playlistName: string) {
    if (this.data.func === 'add') {
      this.data.list.forEach(playlist => {
        if (playlist.playlist_name === playlistName) {
          playlist.songs.push(this.data.song)
          this.dialogRef.close(playlist)
        }
      })
    } else {
      this.data.list.forEach(playlist => {
        if (playlist.playlist_name === playlistName) {
          let temp = playlist.songs.filter(n => n.title !== this.data.song.title)
          this.dialogRef.close(temp)
        }
      })
    }
  }

}
