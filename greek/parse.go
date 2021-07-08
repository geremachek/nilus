package greek

func Parse(raw string, keyb bool) (greek string) {
	var (
		gramma = newGramma(' ', Unaccented, Unmarked, Assumed, false, false)
		chars = []rune(raw)
		end = len(chars)-1

		ch rune
	)

	for i := 0; i <= end; i++ {
		ch = chars[i]

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
				var (
					input = string(ch)
					g rune
				)
				
				if !keyb && i < end {
					input = string(chars[i:i+2])
					i += 1
				}

				g = fromLatin(input, keyb)

				if !keyb && len(input) == 2 && g == ' ' {
					g = fromLatin(string(ch), keyb)
					i -= 1
				}

				if g != ' ' {
					gramma.letter = g
					greek += string(gramma.show())

					gramma.strip()
				} else {
					greek += string(ch)
				}
		}
	}

	return
}
