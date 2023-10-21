package main

import "fmt"

//product interface - playlist
//concrete product1 - popmix
//concrete product2 - hiphopmix
//concrete product3 - kpopmix
//factory - spotify


type Playlist interface {
	Play()
}

type PopMix struct{}

func (p PopMix) Play() {
	fmt.Println("playing pop music mix")
}

type HipHopMix struct{}

func (hh HipHopMix) Play() {
	fmt.Println("playing hiphop music mix")
}

type KPopMix struct{}

func (kp KPopMix) Play() {
	fmt.Println("playing k-pop music mix")
}

type Spotify struct{}

func (s Spotify) PlayMix(key string) Playlist{
	switch key {
	case "pop":
		return PopMix{}
	case "hiphop":
		return HipHopMix{}
	case "kpop":
		return KPopMix{}
	default:
		fmt.Println("Unknown music genre")
		return nil
	}
}

func main() {
	player := Spotify{}

	popMusic := player.PlayMix("pop")
	rockMusic := player.PlayMix("hiphop")
	kpopMusic := player.PlayMix("kpop")

	popMusic.Play()
	rockMusic.Play()
	kpopMusic.Play()
}

