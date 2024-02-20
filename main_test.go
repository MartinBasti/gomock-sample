package main

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/MartinBasti/gomock-sample/makers"
)

// boring sandwich
func TestDefaultSandwich(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := makers.NewMockSandwichMakerInterface(ctrl)

	m.EXPECT().Make(gomock.Eq("something")).Times(1).Return("Mocked something\n")

	MakeASandwich(false, false, m)
}

// one sandwich different filling
func TestHappySandwich(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := makers.NewMockSandwichMakerInterface(ctrl)

	m.EXPECT().Make(gomock.Eq("tuna")).Times(1).Return("Mocked tuna\n")

	MakeASandwich(true, false, m)
}

// two sandwiches
func TestGymSandwich(t *testing.T) {

	ctrl := gomock.NewController(t)
	m := makers.NewMockSandwichMakerInterface(ctrl)

	m.EXPECT().Make(gomock.Eq("tuna")).Times(1).Return("Mocked tuna\n")
	m.EXPECT().Make(gomock.Eq("cheese")).Times(1).Return("Mocked cheese\n")

	MakeASandwich(true, true, m)
}

// sandwich with custom value only for test
func TestMockedSandwich(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := makers.NewMockSandwichMakerInterface(ctrl)

	m.EXPECT().Make(gomock.Any()).AnyTimes().Return("Mocked we don't care sandwich\n")
	MakeASandwich(false, false, m)
}
