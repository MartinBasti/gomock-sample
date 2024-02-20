// SPDX-License-Identifier: Apache-2.0

package makers

//go:generate mockgen -destination mock_sandwich.go -package makers github.com/MartinBasti/gomock-sample/makers SandwichMakerInterface
type SandwichMakerInterface interface {
	Make(string) string
}

type SandwichMaker struct{}

func (sm *SandwichMaker) Make(filling string) string {
	var s string = `
=======
+++++++
=======
`
	if filling == "tuna" {
		s = `
=======
@@@@@@@
=======
`
	}

	if filling == "cheese" {
		s = `
=======
~~~~~~~
=======
`
	}
	return s

}
