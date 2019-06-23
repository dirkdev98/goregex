package match_gen

// QuestLiteral
// Generated regex matcher for the following regex
// ab?
func QuestLiteral(src string) bool {
	idx := 0
	isMatching := true

	questBoolResult22 := false

	staticStr21 := "a"
	staticStr21Len := len(staticStr21)

	if staticStr21Len+idx > len(src) {
		isMatching = false
	} else if staticStr21 == src[idx:idx+staticStr21Len] {
		idx += staticStr21Len
		isMatching = true
	} else {
		isMatching = false
	}
	if isMatching {
		questIdxResult23 := idx

		staticStr24 := "b"
		staticStr24Len := len(staticStr24)

		if staticStr24Len+questIdxResult23 > len(src) {
			questBoolResult22 = false
		} else if staticStr24 == src[questIdxResult23:questIdxResult23+staticStr24Len] {
			questIdxResult23 += staticStr24Len
			questBoolResult22 = true
		} else {
			questBoolResult22 = false
		}

		if questBoolResult22 {
			idx = questIdxResult23
		}

	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of QuestLiteral
}

// QuestDoubleLiteral
// Generated regex matcher for the following regex
// ab?c?
func QuestDoubleLiteral(src string) bool {
	idx := 0
	isMatching := true

	questBoolResult26 := false
	questBoolResult29 := false

	staticStr25 := "a"
	staticStr25Len := len(staticStr25)

	if staticStr25Len+idx > len(src) {
		isMatching = false
	} else if staticStr25 == src[idx:idx+staticStr25Len] {
		idx += staticStr25Len
		isMatching = true
	} else {
		isMatching = false
	}
	if isMatching {
		questIdxResult27 := idx

		staticStr28 := "b"
		staticStr28Len := len(staticStr28)

		if staticStr28Len+questIdxResult27 > len(src) {
			questBoolResult26 = false
		} else if staticStr28 == src[questIdxResult27:questIdxResult27+staticStr28Len] {
			questIdxResult27 += staticStr28Len
			questBoolResult26 = true
		} else {
			questBoolResult26 = false
		}

		if questBoolResult26 {
			idx = questIdxResult27
		}

	}
	if isMatching {
		questIdxResult30 := idx

		staticStr31 := "c"
		staticStr31Len := len(staticStr31)

		if staticStr31Len+questIdxResult30 > len(src) {
			questBoolResult29 = false
		} else if staticStr31 == src[questIdxResult30:questIdxResult30+staticStr31Len] {
			questIdxResult30 += staticStr31Len
			questBoolResult29 = true
		} else {
			questBoolResult29 = false
		}

		if questBoolResult29 {
			idx = questIdxResult30
		}

	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of QuestDoubleLiteral
}

// QuestSingleQuest
// Generated regex matcher for the following regex
// a?
func QuestSingleQuest(src string) bool {
	idx := 0
	isMatching := true

	questBoolResult32 := false
	questIdxResult33 := idx

	staticStr34 := "a"
	staticStr34Len := len(staticStr34)

	if staticStr34Len+questIdxResult33 > len(src) {
		questBoolResult32 = false
	} else if staticStr34 == src[questIdxResult33:questIdxResult33+staticStr34Len] {
		questIdxResult33 += staticStr34Len
		questBoolResult32 = true
	} else {
		questBoolResult32 = false
	}

	if questBoolResult32 {
		idx = questIdxResult33
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of QuestSingleQuest
}
