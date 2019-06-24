package common_gen

// DutchZip
// Generated regex matcher for the following regex
// (\d{4}\s?[a-zA-Z]{2})+
func DutchZip(src string) bool {
	idx := 0
	isMatching := true

	repeatBoolResult81 := false
	questBoolResult84 := false
	repeatBoolResult88 := false
	plusBoolResult91 := false
	repeatBoolResult94 := false
	questBoolResult97 := false
	repeatBoolResult101 := false
	repeatIdxResult82 := idx

	for repeatIdx80 := 1; repeatIdx80 <= 4; repeatIdx80++ {

		if len(src) <= repeatIdxResult82 {
			repeatBoolResult81 = false
		} else {
			charClassRune83 := src[repeatIdxResult82]
			if charClassRune83 >= '0' && charClassRune83 <= '9' {
				repeatIdxResult82 = repeatIdxResult82 + 1
				repeatBoolResult81 = true
			} else {
				repeatBoolResult81 = false
			}
		}

		if !repeatBoolResult81 {
			if repeatIdx80 > 4 {
				idx = repeatIdxResult82
				isMatching = true
				break
			} else {
				isMatching = false
				break
			}
		} else if repeatBoolResult81 && repeatIdx80 == 4 {
			idx = repeatIdxResult82
			isMatching = true
		}
	}
	if isMatching {
		questIdxResult85 := idx

		if len(src) <= questIdxResult85 {
			questBoolResult84 = false
		} else {
			charClassRune86 := src[questIdxResult85]
			if (charClassRune86 >= '\t' && charClassRune86 <= '\n') || (charClassRune86 >= '\f' && charClassRune86 <= '\r') || (charClassRune86 == ' ') {
				questIdxResult85 = questIdxResult85 + 1
				questBoolResult84 = true
			} else {
				questBoolResult84 = false
			}
		}

		if questBoolResult84 {
			idx = questIdxResult85
		}

	}
	if isMatching {
		repeatIdxResult89 := idx

		for repeatIdx87 := 1; repeatIdx87 <= 2; repeatIdx87++ {

			if len(src) <= repeatIdxResult89 {
				repeatBoolResult88 = false
			} else {
				charClassRune90 := src[repeatIdxResult89]
				if (charClassRune90 >= 'A' && charClassRune90 <= 'Z') || (charClassRune90 >= 'a' && charClassRune90 <= 'z') {
					repeatIdxResult89 = repeatIdxResult89 + 1
					repeatBoolResult88 = true
				} else {
					repeatBoolResult88 = false
				}
			}

			if !repeatBoolResult88 {
				if repeatIdx87 > 2 {
					idx = repeatIdxResult89
					isMatching = true
					break
				} else {
					isMatching = false
					break
				}
			} else if repeatBoolResult88 && repeatIdx87 == 2 {
				idx = repeatIdxResult89
				isMatching = true
			}
		}

	}
	plusIdxResult92 := idx

	if isMatching {
		for {
			repeatIdxResult95 := plusIdxResult92

			for repeatIdx93 := 1; repeatIdx93 <= 4; repeatIdx93++ {

				if len(src) <= repeatIdxResult95 {
					repeatBoolResult94 = false
				} else {
					charClassRune96 := src[repeatIdxResult95]
					if charClassRune96 >= '0' && charClassRune96 <= '9' {
						repeatIdxResult95 = repeatIdxResult95 + 1
						repeatBoolResult94 = true
					} else {
						repeatBoolResult94 = false
					}
				}

				if !repeatBoolResult94 {
					if repeatIdx93 > 4 {
						plusIdxResult92 = repeatIdxResult95
						plusBoolResult91 = true
						break
					} else {
						plusBoolResult91 = false
						break
					}
				} else if repeatBoolResult94 && repeatIdx93 == 4 {
					plusIdxResult92 = repeatIdxResult95
					plusBoolResult91 = true
				}
			}
			if plusBoolResult91 {
				questIdxResult98 := plusIdxResult92

				if len(src) <= questIdxResult98 {
					questBoolResult97 = false
				} else {
					charClassRune99 := src[questIdxResult98]
					if (charClassRune99 >= '\t' && charClassRune99 <= '\n') || (charClassRune99 >= '\f' && charClassRune99 <= '\r') || (charClassRune99 == ' ') {
						questIdxResult98 = questIdxResult98 + 1
						questBoolResult97 = true
					} else {
						questBoolResult97 = false
					}
				}

				if questBoolResult97 {
					plusIdxResult92 = questIdxResult98
				}

			}
			if plusBoolResult91 {
				repeatIdxResult102 := plusIdxResult92

				for repeatIdx100 := 1; repeatIdx100 <= 2; repeatIdx100++ {

					if len(src) <= repeatIdxResult102 {
						repeatBoolResult101 = false
					} else {
						charClassRune103 := src[repeatIdxResult102]
						if (charClassRune103 >= 'A' && charClassRune103 <= 'Z') || (charClassRune103 >= 'a' && charClassRune103 <= 'z') {
							repeatIdxResult102 = repeatIdxResult102 + 1
							repeatBoolResult101 = true
						} else {
							repeatBoolResult101 = false
						}
					}

					if !repeatBoolResult101 {
						if repeatIdx100 > 2 {
							plusIdxResult92 = repeatIdxResult102
							plusBoolResult91 = true
							break
						} else {
							plusBoolResult91 = false
							break
						}
					} else if repeatBoolResult101 && repeatIdx100 == 2 {
						plusIdxResult92 = repeatIdxResult102
						plusBoolResult91 = true
					}
				}

			}

			if !plusBoolResult91 {
				break
			} else {
				isMatching = plusBoolResult91
				idx = plusIdxResult92
			}
		}
	}

	if idx != len(src) {
		isMatching = false
	}

	return isMatching
	// End of DutchZip
}
