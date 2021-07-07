package greek

func Parse(raw string) (greek string) {
	gramma := newGramma(' ', Unaccented, Unmarked, Assumed, false, false)

	for _, ch := range []rune(raw) {
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
			
			case '?':  greek += ";"
			case ';':  greek += "·"
		
			default:
				if ((ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')) && (ch != 'q' && ch != 'Q') {
					gramma.letter = fromLatin(ch)
					greek += string(gramma.show())

					gramma.strip()
				} else {
					greek += string(ch)
				}
		}
	}

	return
}
