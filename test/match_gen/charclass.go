package match_gen

// CharClassSingle
// Generated regex matcher for the following regex
// [abc]
func CharClassSingle(src string) bool {
	idx := 0
	isMatching := true

	if len(src) <= idx {
		isMatching = false
	} else {
		charClassRune74 := src[idx]
		if charClassRune74 >= 'a' && charClassRune74 <= 'c' {
			idx = idx + 1
			isMatching = true
		} else {
			isMatching = false
		}
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of CharClassSingle
}

// CharClassRange
// Generated regex matcher for the following regex
// [\w]
func CharClassRange(src string) bool {
	idx := 0
	isMatching := true

	if len(src) <= idx {
		isMatching = false
	} else {
		charClassRune75 := src[idx]
		if (charClassRune75 >= '0' && charClassRune75 <= '9') || (charClassRune75 >= 'A' && charClassRune75 <= 'Z') || (charClassRune75 == '_') || (charClassRune75 >= 'a' && charClassRune75 <= 'z') {
			idx = idx + 1
			isMatching = true
		} else {
			isMatching = false
		}
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of CharClassRange
}

// CharClassPlusRange
// Generated regex matcher for the following regex
// [\d]+
func CharClassPlusRange(src string) bool {
	idx := 0
	isMatching := true

	plusBoolResult77 := false

	if len(src) <= idx {
		isMatching = false
	} else {
		charClassRune76 := src[idx]
		if charClassRune76 >= '0' && charClassRune76 <= '9' {
			idx = idx + 1
			isMatching = true
		} else {
			isMatching = false
		}
	}
	plusIdxResult78 := idx

	if isMatching {
		for {

			if len(src) <= plusIdxResult78 {
				plusBoolResult77 = false
			} else {
				charClassRune79 := src[plusIdxResult78]
				if charClassRune79 >= '0' && charClassRune79 <= '9' {
					plusIdxResult78 = plusIdxResult78 + 1
					plusBoolResult77 = true
				} else {
					plusBoolResult77 = false
				}
			}

			if !plusBoolResult77 {
				break
			} else {
				isMatching = plusBoolResult77
				idx = plusIdxResult78
			}
		}
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of CharClassPlusRange
}
