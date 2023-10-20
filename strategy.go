package main

import "fmt"

//avatar aang enemy encounter

//strategy interface - use bending 
//concrete strategy 1 - use air to float the enemy
//concrete strategy 2 - use water to freeze enemy
//concrete strategy 3 - use earth to immobilize the enemy
//concrete strategy 4 - use fire to kill the enemy
//context - Avatar Aang

type BendingStrategy interface {
	UseBending()
}

type AirBendingStrategy struct{}

func (a *AirBendingStrategy) UseBending() {
	fmt.Printf("Aand made the enemy float into the air and fall.\n")
}

type WaterBendingStrategy struct{}

func (w *WaterBendingStrategy) UseBending() {
	fmt.Printf("Aang frozen the enemy.\n")
}

type EarthBendingStrategy struct{}

func (e *EarthBendingStrategy) UseBending() {
	fmt.Printf("Aand captured the enemy inside the earth trap.\n")
}

type FireBendingStrategy struct{}

func (f *FireBendingStrategy) UseBending() {
	fmt.Printf("Aang damaged the enemy using fireblast.\n")
}

type Aang struct {
	bendingStrategy BendingStrategy
}

func (a *Aang) EncounterEnemy(bendingStrategy BendingStrategy) {
	fmt.Printf("Enemy appears in front of the Avatar:\n")
	bendingStrategy.UseBending()
}

func main() {
	airStrategy := &AirBendingStrategy{}
	waterStrategy := &WaterBendingStrategy{}
	earthStrategy := &EarthBendingStrategy{}
	fireStrategy := &FireBendingStrategy{}

	aang := &Aang{}

	aang.EncounterEnemy(airStrategy)
	aang.EncounterEnemy(waterStrategy)
	aang.EncounterEnemy(earthStrategy)
	aang.EncounterEnemy(fireStrategy)
}