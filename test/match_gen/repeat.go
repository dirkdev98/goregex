package match_gen

// RepeatLiteral
// Generated regex matcher for the following regex
// a{2}
func RepeatLiteral(src string) bool {
	idx := 0
	isMatching := true

	repeatBoolResult36 := false
	repeatIdxResult37 := idx

	for repeatIdx35 := 1; repeatIdx35 <= 2; repeatIdx35++ {

		staticStr38 := "a"
		staticStr38Len := len(staticStr38)

		if staticStr38Len+repeatIdxResult37 > len(src) {
			repeatBoolResult36 = false
		} else if staticStr38 == src[repeatIdxResult37:repeatIdxResult37+staticStr38Len] {
			repeatIdxResult37 += staticStr38Len
			repeatBoolResult36 = true
		} else {
			repeatBoolResult36 = false
		}

		if !repeatBoolResult36 {
			if repeatIdx35 > 2 {
				idx = repeatIdxResult37
				break
			} else {
				isMatching = false
				break
			}
		} else if repeatBoolResult36 && repeatIdx35 == 2 {
			idx = repeatIdxResult37
		}
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of RepeatLiteral
}

// RepeatLiteralMinMax
// Generated regex matcher for the following regex
// a{2,5}
func RepeatLiteralMinMax(src string) bool {
	idx := 0
	isMatching := true

	repeatBoolResult40 := false
	repeatIdxResult41 := idx

	for repeatIdx39 := 1; repeatIdx39 <= 5; repeatIdx39++ {

		staticStr42 := "a"
		staticStr42Len := len(staticStr42)

		if staticStr42Len+repeatIdxResult41 > len(src) {
			repeatBoolResult40 = false
		} else if staticStr42 == src[repeatIdxResult41:repeatIdxResult41+staticStr42Len] {
			repeatIdxResult41 += staticStr42Len
			repeatBoolResult40 = true
		} else {
			repeatBoolResult40 = false
		}

		if !repeatBoolResult40 {
			if repeatIdx39 > 2 {
				idx = repeatIdxResult41
				break
			} else {
				isMatching = false
				break
			}
		} else if repeatBoolResult40 && repeatIdx39 == 5 {
			idx = repeatIdxResult41
		}
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of RepeatLiteralMinMax
}

// RepeatWithAlternate
// Generated regex matcher for the following regex
// ab|ac{2}
func RepeatWithAlternate(src string) bool {
	idx := 0
	isMatching := true

	hasSet45 := false
	alternateResult46 := false
	alternateResult48 := false
	repeatBoolResult50 := false

	staticStr43 := "a"
	staticStr43Len := len(staticStr43)

	if staticStr43Len+idx > len(src) {
		isMatching = false
	} else if staticStr43 == src[idx:idx+staticStr43Len] {
		idx += staticStr43Len
		isMatching = true
	} else {
		isMatching = false
	}
	if isMatching {

		for i44 := 1; i44 <= 2; i44++ {

			if i44 == 1 {

				staticStr47 := "b"
				staticStr47Len := len(staticStr47)

				if staticStr47Len+idx > len(src) {
					alternateResult46 = false
				} else if staticStr47 == src[idx:idx+staticStr47Len] {
					idx += staticStr47Len
					alternateResult46 = true
				} else {
					alternateResult46 = false
				}

				if alternateResult46 {
					isMatching = true
					hasSet45 = true
					break
				}
			}

			if i44 == 2 {
				repeatIdxResult51 := idx

				for repeatIdx49 := 1; repeatIdx49 <= 2; repeatIdx49++ {

					staticStr52 := "c"
					staticStr52Len := len(staticStr52)

					if staticStr52Len+repeatIdxResult51 > len(src) {
						repeatBoolResult50 = false
					} else if staticStr52 == src[repeatIdxResult51:repeatIdxResult51+staticStr52Len] {
						repeatIdxResult51 += staticStr52Len
						repeatBoolResult50 = true
					} else {
						repeatBoolResult50 = false
					}

					if !repeatBoolResult50 {
						if repeatIdx49 > 2 {
							idx = repeatIdxResult51
							break
						} else {
							alternateResult48 = false
							break
						}
					} else if repeatBoolResult50 && repeatIdx49 == 2 {
						idx = repeatIdxResult51
					}
				}

				if alternateResult48 {
					isMatching = true
					hasSet45 = true
					break
				}
			}

		}
		if !hasSet45 {
			isMatching = false
		}

	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of RepeatWithAlternate
}
