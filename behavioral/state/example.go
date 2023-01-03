package state

import "fmt"

// State
type State interface {
	clickLock()
	clickPlay()
	clickNext()
	clickPrevious()
}

// Context
type AudioPlayer struct {
	State       State
	UI          string
	Volumn      int
	Playlist    []string
	CurrentSong int
}

var _ State = &AudioPlayer{} // Context can implement the State Interface.

func (a *AudioPlayer) Init() {
	a.State = &ReadyState{}
}

func (a *AudioPlayer) ChangeState(state State) {
	a.State = state
}

func (a *AudioPlayer) clickLock() {
	a.State.clickLock()
}

func (a *AudioPlayer) clickPlay() {
	a.State.clickPlay()
}

func (a *AudioPlayer) clickNext() {
	a.State.clickNext()
}

func (a *AudioPlayer) clickPrevious() {
	a.State.clickPrevious()
}

func (a *AudioPlayer) playMusic() {
	fmt.Println("playing", a.Playlist[a.CurrentSong])
}

// Concrete States
type LockedState struct {
	player *AudioPlayer // Backreference to the context, so concrete states can transition the context to another state.
}

type ReadyState struct {
	player *AudioPlayer
}

type PlayingState struct {
	player *AudioPlayer
}

var _ State = &LockedState{}
var _ State = &ReadyState{}
var _ State = &PlayingState{}

func (l *LockedState) clickLock() {
	fmt.Println("player unlocked")
	l.player.ChangeState(&ReadyState{})
}

func (l *LockedState) clickPlay() {
	fmt.Println("play unlock first")
}

func (l *LockedState) clickNext() {}

func (l *LockedState) clickPrevious() {}

func (r *ReadyState) clickLock() {
	r.player.ChangeState(&LockedState{})
}

func (r *ReadyState) clickPlay() {
	r.player.playMusic()
	r.player.ChangeState(&PlayingState{})
}

func (r *ReadyState) clickNext() {
	r.player.CurrentSong = (r.player.CurrentSong + 1) % len(r.player.Playlist)
	r.player.ChangeState(&PlayingState{})
}

func (r *ReadyState) clickPrevious() {
	r.player.CurrentSong = (r.player.CurrentSong + len(r.player.Playlist) - 1) % len(r.player.Playlist)
	r.player.ChangeState(&PlayingState{})
}

func (p *PlayingState) clickLock() {
	p.player.ChangeState(&LockedState{})
}

func (p *PlayingState) clickPlay() {
	p.player.playMusic()
}

func (p *PlayingState) clickNext() {
	p.player.CurrentSong = (p.player.CurrentSong + 1) % len(p.player.Playlist)
	p.player.playMusic()
}

func (p *PlayingState) clickPrevious() {
	p.player.CurrentSong = (p.player.CurrentSong - 1 + len(p.player.Playlist)) % len(p.player.Playlist)
	p.player.playMusic()
}
