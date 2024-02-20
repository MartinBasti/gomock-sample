// SPDX-License-Identifier: Apache-2.0

package makers

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
