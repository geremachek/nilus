package main

// parse the "markup" language

func parse(raw string, keyb bool) (greek string) {
	var (
		gramma = NewGramma(0, Unaccented, Unmarked, Assumed, false, false)
		chars = []rune(raw)
		end = len(chars)-1

		ch rune
	)

	for i := 0; i <= end; i++ {
		ch = chars[i]

		switch ch {

			// convert symbols into diacritics

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

				// deal with digraphs and alternate characters...

				if !keyb && i < end {
					input = string(chars[i:i+2])

					g = fromLatin(input, keyb)

					// if it fails, deal with the first character, otherwise increment the index

					if g == 0 {
						g = fromLatin(string(ch), keyb)
					} else {
						i += 1
					}
				} else {
					g = fromLatin(input, keyb)
				}

				// add to the output string

				if g != 0 {
					// ending sigma converted when it appears at the end of a line or encounters puncuation

					if !keyb && g == Sigma && (i == end || isSimplePunct(chars[i+1])) {
						g = UltimateSigma
					}

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

// check for the punctuation we need...

func isSimplePunct(ch rune) bool {
	switch ch {
		case ' ':
		case '.':
		case '!':
		case '?':
		case ';':
		case '"':
		default: return false	
	}

	return true
}