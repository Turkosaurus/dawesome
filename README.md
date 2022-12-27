# dawesome

## Repo Structure
```
.
├── main.go
└── /cmd
    └── dawesome.go
```
## Purpose
Create a live, semi-asyncronous musical looper for use with students or collaborators.

## Overview

### Initial Server Response
- accepts initial request
- sends client and GUI

### Client Loading
- loads console page
    - connection status
    - microphone permissions
    - boot local server?
    - establish streaming handshake with server
    - tempo established, tied to UTC, sent to server
- establish peer connection, update in connection status

## Shared Count
- shared global tempo syncronized across clients

## Client Sending of Recorded Media
- record locally and store to buffer and/or local file
- package data in chunks and send to server

## Client Playback of Received Media
- checks for new update perdioically
- playback received media on next avaiable measure from global tempo