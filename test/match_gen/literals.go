package match_gen

// LiteralSimple
// Generated regex matcher for the following regex
// literal
func LiteralSimple(src string) bool {
	idx := 0
	isMatching := true

	staticStr0 := "literal"
	staticStr0Len := len(staticStr0)

	if staticStr0Len+idx > len(src) {
		isMatching = false
	} else if staticStr0 == src[idx:idx+staticStr0Len] {
		idx += staticStr0Len
		isMatching = true
	} else {
		isMatching = false
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of LiteralSimple
}
