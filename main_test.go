package main

import (
	"testing"

	"github.com/MartinBasti/gomock-sample/makers"
)

// boring sandwich
func TestDefaultSandwich(t *testing.T) {
	MakeASandwich(false, false, &makers.SandwichMaker{})
}

// one sandwich different filling
func TestHappySandwich(t *testing.T) {
	MakeASandwich(true, false, &makers.SandwichMaker{})
}

// two sandwiches
func TestGymSandwich(t *testing.T) {
	MakeASandwich(false, true, &makers.SandwichMaker{})
}

// sandwich with custom value only for test
func TestMockedSandwich(t *testing.T) {
	MakeASandwich(false, false, &makers.SandwichMaker{})
}
