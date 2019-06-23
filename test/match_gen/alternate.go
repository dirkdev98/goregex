package match_gen

// AlternateLiterals
// Generated regex matcher for the following regex
// ab|cd
func AlternateLiterals(src string) bool {
	idx := 0
	isMatching := true

	hasSet2 := false
	alternateResult3 := false
	alternateResult5 := false

	for i1 := 1; i1 <= 2; i1++ {

		if i1 == 1 {

			staticStr4 := "ab"
			staticStr4Len := len(staticStr4)

			if staticStr4Len+idx > len(src) {
				alternateResult3 = false
			} else if staticStr4 == src[idx:idx+staticStr4Len] {
				idx += staticStr4Len
				alternateResult3 = true
			} else {
				alternateResult3 = false
			}

			if alternateResult3 {
				isMatching = true
				hasSet2 = true
				break
			}
		}

		if i1 == 2 {

			staticStr6 := "cd"
			staticStr6Len := len(staticStr6)

			if staticStr6Len+idx > len(src) {
				alternateResult5 = false
			} else if staticStr6 == src[idx:idx+staticStr6Len] {
				idx += staticStr6Len
				alternateResult5 = true
			} else {
				alternateResult5 = false
			}

			if alternateResult5 {
				isMatching = true
				hasSet2 = true
				break
			}
		}

	}
	if !hasSet2 {
		isMatching = false
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of AlternateLiterals
}

// AlternateMultipleLiterals
// Generated regex matcher for the following regex
// ab|cd|ef|gh|ij|kl
func AlternateMultipleLiterals(src string) bool {
	idx := 0
	isMatching := true

	hasSet8 := false
	alternateResult9 := false
	alternateResult11 := false
	alternateResult13 := false
	alternateResult15 := false
	alternateResult17 := false
	alternateResult19 := false

	for i7 := 1; i7 <= 6; i7++ {

		if i7 == 1 {

			staticStr10 := "ab"
			staticStr10Len := len(staticStr10)

			if staticStr10Len+idx > len(src) {
				alternateResult9 = false
			} else if staticStr10 == src[idx:idx+staticStr10Len] {
				idx += staticStr10Len
				alternateResult9 = true
			} else {
				alternateResult9 = false
			}

			if alternateResult9 {
				isMatching = true
				hasSet8 = true
				break
			}
		}

		if i7 == 2 {

			staticStr12 := "cd"
			staticStr12Len := len(staticStr12)

			if staticStr12Len+idx > len(src) {
				alternateResult11 = false
			} else if staticStr12 == src[idx:idx+staticStr12Len] {
				idx += staticStr12Len
				alternateResult11 = true
			} else {
				alternateResult11 = false
			}

			if alternateResult11 {
				isMatching = true
				hasSet8 = true
				break
			}
		}

		if i7 == 3 {

			staticStr14 := "ef"
			staticStr14Len := len(staticStr14)

			if staticStr14Len+idx > len(src) {
				alternateResult13 = false
			} else if staticStr14 == src[idx:idx+staticStr14Len] {
				idx += staticStr14Len
				alternateResult13 = true
			} else {
				alternateResult13 = false
			}

			if alternateResult13 {
				isMatching = true
				hasSet8 = true
				break
			}
		}

		if i7 == 4 {

			staticStr16 := "gh"
			staticStr16Len := len(staticStr16)

			if staticStr16Len+idx > len(src) {
				alternateResult15 = false
			} else if staticStr16 == src[idx:idx+staticStr16Len] {
				idx += staticStr16Len
				alternateResult15 = true
			} else {
				alternateResult15 = false
			}

			if alternateResult15 {
				isMatching = true
				hasSet8 = true
				break
			}
		}

		if i7 == 5 {

			staticStr18 := "ij"
			staticStr18Len := len(staticStr18)

			if staticStr18Len+idx > len(src) {
				alternateResult17 = false
			} else if staticStr18 == src[idx:idx+staticStr18Len] {
				idx += staticStr18Len
				alternateResult17 = true
			} else {
				alternateResult17 = false
			}

			if alternateResult17 {
				isMatching = true
				hasSet8 = true
				break
			}
		}

		if i7 == 6 {

			staticStr20 := "kl"
			staticStr20Len := len(staticStr20)

			if staticStr20Len+idx > len(src) {
				alternateResult19 = false
			} else if staticStr20 == src[idx:idx+staticStr20Len] {
				idx += staticStr20Len
				alternateResult19 = true
			} else {
				alternateResult19 = false
			}

			if alternateResult19 {
				isMatching = true
				hasSet8 = true
				break
			}
		}

	}
	if !hasSet8 {
		isMatching = false
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of AlternateMultipleLiterals
}
