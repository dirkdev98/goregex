package match_gen

// StarLiteral
// Generated regex matcher for the following regex
// a*
func StarLiteral(src string) bool {
	idx := 0
	isMatching := true

	starBoolResult65 := false
	starIdxResult66 := idx

	for {

		staticStr67 := "a"
		staticStr67Len := len(staticStr67)

		if staticStr67Len+starIdxResult66 > len(src) {
			starBoolResult65 = false
		} else if staticStr67 == src[starIdxResult66:starIdxResult66+staticStr67Len] {
			starIdxResult66 += staticStr67Len
			starBoolResult65 = true
		} else {
			starBoolResult65 = false
		}

		if !starBoolResult65 {
			break
		} else {
			isMatching = starBoolResult65
			idx = starIdxResult66
		}
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of StarLiteral
}

// StarConcatLiteral
// Generated regex matcher for the following regex
// a*b*
func StarConcatLiteral(src string) bool {
	idx := 0
	isMatching := true

	starBoolResult68 := false
	starBoolResult71 := false
	starIdxResult69 := idx

	for {

		staticStr70 := "a"
		staticStr70Len := len(staticStr70)

		if staticStr70Len+starIdxResult69 > len(src) {
			starBoolResult68 = false
		} else if staticStr70 == src[starIdxResult69:starIdxResult69+staticStr70Len] {
			starIdxResult69 += staticStr70Len
			starBoolResult68 = true
		} else {
			starBoolResult68 = false
		}

		if !starBoolResult68 {
			break
		} else {
			isMatching = starBoolResult68
			idx = starIdxResult69
		}
	}
	if isMatching {
		starIdxResult72 := idx

		for {

			staticStr73 := "b"
			staticStr73Len := len(staticStr73)

			if staticStr73Len+starIdxResult72 > len(src) {
				starBoolResult71 = false
			} else if staticStr73 == src[starIdxResult72:starIdxResult72+staticStr73Len] {
				starIdxResult72 += staticStr73Len
				starBoolResult71 = true
			} else {
				starBoolResult71 = false
			}

			if !starBoolResult71 {
				break
			} else {
				isMatching = starBoolResult71
				idx = starIdxResult72
			}
		}

	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of StarConcatLiteral
}
