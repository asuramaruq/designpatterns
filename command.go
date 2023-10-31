package main

import (
	"fmt"
)

type Command interface {
	Execute()
}

type SpeakerInterface interface {
	Play()
	Pause()
}

type Speaker struct {
	IsPlaying bool
}

func (s *Speaker) Play() {
	if !s.IsPlaying {
		fmt.Printf("playing music\n")
		s.IsPlaying = true
	} else {
		fmt.Printf("music is already playing\n")
	}
}

func (s *Speaker) Pause() {
	if s.IsPlaying {
		fmt.Printf("music paused\n")
		s.IsPlaying = false
	} else {
		fmt.Printf("music is already paused\n")
	}
}

type PlayCommand struct {
	Speaker SpeakerInterface
}

func (plc *PlayCommand) Execute() {
	plc.Speaker.Play()
}

type PauseCommand struct {
	Speaker SpeakerInterface
}

func (psc *PauseCommand) Execute() {
	psc.Speaker.Pause()
}

type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.Execute()
}

func main() {
	speaker := &Speaker{}

	playCommand := &PlayCommand{Speaker: speaker}
	pauseCommand := &PauseCommand{Speaker: speaker}

	PlayButton := &Button{Command: playCommand}
	PauseButton := &Button{Command: pauseCommand}
	PlayRemote := &Button{Command: playCommand}
	PauseRemote := &Button{Command: pauseCommand}

	PlayButton.Press()
	PlayRemote.Press()
	PauseRemote.Press()
	PauseButton.Press()
}
