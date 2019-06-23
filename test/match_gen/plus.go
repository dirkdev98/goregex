package match_gen

// PlusLiteral
// Generated regex matcher for the following regex
// a+
func PlusLiteral(src string) bool {
	idx := 0
	isMatching := true

	plusBoolResult54 := false

	staticStr53 := "a"
	staticStr53Len := len(staticStr53)

	if staticStr53Len+idx > len(src) {
		isMatching = false
	} else if staticStr53 == src[idx:idx+staticStr53Len] {
		idx += staticStr53Len
		isMatching = true
	} else {
		isMatching = false
	}
	plusIdxResult55 := idx

	if isMatching {
		for {

			staticStr56 := "a"
			staticStr56Len := len(staticStr56)

			if staticStr56Len+plusIdxResult55 > len(src) {
				plusBoolResult54 = false
			} else if staticStr56 == src[plusIdxResult55:plusIdxResult55+staticStr56Len] {
				plusIdxResult55 += staticStr56Len
				plusBoolResult54 = true
			} else {
				plusBoolResult54 = false
			}

			if !plusBoolResult54 {
				break
			} else {
				isMatching = plusBoolResult54
				idx = plusIdxResult55
			}
		}
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of PlusLiteral
}

// PlusConcatLiteral
// Generated regex matcher for the following regex
// a+b+
func PlusConcatLiteral(src string) bool {
	idx := 0
	isMatching := true

	plusBoolResult58 := false
	plusBoolResult62 := false

	staticStr57 := "a"
	staticStr57Len := len(staticStr57)

	if staticStr57Len+idx > len(src) {
		isMatching = false
	} else if staticStr57 == src[idx:idx+staticStr57Len] {
		idx += staticStr57Len
		isMatching = true
	} else {
		isMatching = false
	}
	plusIdxResult59 := idx

	if isMatching {
		for {

			staticStr60 := "a"
			staticStr60Len := len(staticStr60)

			if staticStr60Len+plusIdxResult59 > len(src) {
				plusBoolResult58 = false
			} else if staticStr60 == src[plusIdxResult59:plusIdxResult59+staticStr60Len] {
				plusIdxResult59 += staticStr60Len
				plusBoolResult58 = true
			} else {
				plusBoolResult58 = false
			}

			if !plusBoolResult58 {
				break
			} else {
				isMatching = plusBoolResult58
				idx = plusIdxResult59
			}
		}
	}
	if isMatching {

		staticStr61 := "b"
		staticStr61Len := len(staticStr61)

		if staticStr61Len+idx > len(src) {
			isMatching = false
		} else if staticStr61 == src[idx:idx+staticStr61Len] {
			idx += staticStr61Len
			isMatching = true
		} else {
			isMatching = false
		}
		plusIdxResult63 := idx

		if isMatching {
			for {

				staticStr64 := "b"
				staticStr64Len := len(staticStr64)

				if staticStr64Len+plusIdxResult63 > len(src) {
					plusBoolResult62 = false
				} else if staticStr64 == src[plusIdxResult63:plusIdxResult63+staticStr64Len] {
					plusIdxResult63 += staticStr64Len
					plusBoolResult62 = true
				} else {
					plusBoolResult62 = false
				}

				if !plusBoolResult62 {
					break
				} else {
					isMatching = plusBoolResult62
					idx = plusIdxResult63
				}
			}
		}

	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of PlusConcatLiteral
}
