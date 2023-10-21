package main

import "fmt"

//component - sandwich
//concrete component - basic sandwich 
//base decorator - toppingdecorator
//concrete decorators - all the toppings themselves

type Sandwich interface {
	Description() string
}

type BasicSandwich struct{}

func (b *BasicSandwich) Description() string {
	return "Slice of Bread"
}

type ToppingDecorator struct {
	sandwich Sandwich
}

func (td *ToppingDecorator) Description() string {
	return td.sandwich.Description()
}

type CheeseTopping struct {
	sandwich Sandwich
}

func (ct *CheeseTopping) Description() string {
	return ct.sandwich.Description() + ", Cheese"
}

type LettuceTopping struct {
	sandwich Sandwich
}

func (lt *LettuceTopping) Description() string {
	return lt.sandwich.Description() + ", Lettuce"
}

type SausageTopping struct {
	sandwich Sandwich
}

func (st *SausageTopping) Description() string {
	return st.sandwich.Description() + ", Sausage"
}

type SauseTopping struct {
	sandwich Sandwich
}

func (st *SauseTopping) Description() string {
	return st.sandwich.Description() + ", Sause"
}

type BreadTopping struct {
	sandwich Sandwich
}

func (bt *BreadTopping) Description() string {
	return bt.sandwich.Description() + ", Slice of Bread"
}

func main() {
	basicSandwich := &BasicSandwich{}

	sandwichWithS := &SauseTopping{
		sandwich: basicSandwich,
	}

	sandwichWithSS := &SausageTopping{
		sandwich: sandwichWithS,
	}

	sandwichWithSSC := &CheeseTopping{
		sandwich: sandwichWithSS,
	}

	sandwichWithSSCL := &LettuceTopping{
		sandwich: sandwichWithSSC,
	}

	sandwichWithSSCLS := &SauseTopping{
		sandwich: sandwichWithSSCL,
	}

	sandwichWithSSCLSB := &BreadTopping{
		sandwich: sandwichWithSSCLS,
	}

	fmt.Println("Sandwich description:", sandwichWithSSCLSB.Description())
}
