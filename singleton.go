package main

import (
	"fmt"
	"sync"
)

type Roblux struct {
}

var instance *Roblux
var once sync.Once

func GetInstance() *Roblux {
	once.Do(func() {
		instance = &Roblux{}
	})
	return instance
}

type Player struct {
	UserID   string
	Username string
	Device   string
}

func (gs *Roblux) Connect(player *Player) {
	fmt.Printf("%s connected to the game server using %s.\n", player.Username, player.Device)
}

func main() {
	roblux := GetInstance()

	phonePlayer := &Player{UserID: "11111", Username: "mhmhmh", Device: "Mobile"}
	pcPlayer := &Player{UserID: "22222", Username: "uzi", Device: "PC"}
	playStationPlayer := &Player{UserID: "33333", Username: "highwaytohell", Device: "PlayStation"}
	xboxPlayer := &Player{UserID: "44444", Username: "whotf", Device: "Xbox"}

	roblux.Connect(phonePlayer)
	roblux.Connect(pcPlayer)
	roblux.Connect(playStationPlayer)
	roblux.Connect(xboxPlayer)
}
