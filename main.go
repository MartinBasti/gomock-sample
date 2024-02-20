// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/MartinBasti/gomock-sample/makers"
)

func MakeASandwich(areYouHappy, didYouExcersise bool, maker makers.SandwichMakerInterface) {
	if areYouHappy {
		fmt.Printf("Happy sandwich:%s", maker.Make("tuna"))
	} else {
		fmt.Printf("Standard sandwich:%s", maker.Make("something"))
	}

	if didYouExcersise {
		fmt.Printf("\nHave an extra cheese sandwich:%s", maker.Make("cheese"))
	}
}

func main() {
	MakeASandwich(true, true, &makers.SandwichMaker{})
}
