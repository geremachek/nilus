package greek

func Parse(raw string, arch bool) (greek string) {
	gramma := newGramma(' ', Unaccented, Unmarked, Assumed, false, false)
	
	for _, ch := range []rune(raw) {
		switch ch {
			case '?':  greek += ";"
			case ';':  greek += "·"
			default:
				if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
					if arch {
						greek += string(fromLatin(ch, true))
					} else {
						gramma.letter = fromLatin(ch, false)
						greek += string(gramma.show())

						gramma.strip()
					}
				} else if !arch {
					switch ch {
						case '\'': gramma.accent    = Acute
						case '`':  gramma.accent    = Grave
						case '~':  gramma.accent    = Circumflex
						case '|':  gramma.accent    = Tonos

						case '[':  gramma.breath    = Rough
						case ']':  gramma.breath    = Smooth

						case '_':  gramma.length    = Long
						case '^':  gramma.length    = Short
			
						case ':':  gramma.diaeresis = true
						case '*':  gramma.iota      = true

						default: greek += string(ch)
					}
				} else {
					greek += string(ch)
				}
		}
	}

	return
}
