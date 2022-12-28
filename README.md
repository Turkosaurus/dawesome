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

## Timeline Behavior
The timeline is bifurcated left (past) and right (future) of the center (global now). Sound bars travel right to left, containing only the visual metronome and beat/bar demarcations when players are silent. When player A begins recording, player B's timeline is frozen until A is complete. B hears A after A is complete, or after a specified offset. The timeline "unfreezes" when recording is complete, and both players (re)listen in exact simultaneous real-time. There will be a 1 bar gap between "call" and "response", but both players experience an uninterrupted, continuous tempo.

## Chunk Management: Three Timelines
1. Own timeine: most recent played measure(s) queued for sending to global playback
1. Collaborator timeline: most recent played measure(s) from collaborator
1. Saved measure(s)

## Sections
Sections A, B, C, etc. can be chained together from past timelines.