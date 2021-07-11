package greek

import "unicode"

// structure that represents an accented character

type Gramma struct {
	letter rune
	accent Accent

	breath Breathing
	length Length

	diaeresis bool
	iota bool
}

func newGramma(let rune, a Accent, b Breathing, l Length, d, i bool) Gramma {
	return Gramma { let, a, b, l, d, i }
}

// remove all diacritics

func (g *Gramma) strip() {
	g.letter = ' '
	g.accent = Unaccented

	g.breath = Unmarked
	g.length = Assumed

	g.diaeresis = false
	g.iota = false
}

// show the character with diagritics

func (g Gramma) show() (shown rune) {
	upper, low := handleCase(g.letter)
	shown = low
	
	// non-standard diacritic application

	if low == Rho {
		switch g.breath {
			case Rough:  shown = 'ῥ'
			case Smooth: shown = 'ῤ'
		}
	} else if low == HookedUpsilon {
		if g.accent == Acute {
			shown += 1
		} else if g.diaeresis {
			shown += 2
		}
	}

	defer func() { if upper { shown = unicode.ToUpper(shown) } }()

	if low == Beta  || low == Gamma         || low == Delta   || low == Zeta ||
	   low == Theta || low == Kappa         || low == Lambda  || low == Mu   ||
	   low == Nu    || low == Xi            || low == Pi      || low == Rho  ||
	   low == Sigma || low == Tau           || low == Phi     || low == Chi  ||
	   low == Psi   || low == UltimateSigma || low == HookedUpsilon {
		return
	}

	if g.breath != Unmarked {
		switch low {
			case Alpha:   shown = 'ἀ'
			case Epsilon: shown = 'ἐ'
			case Eta:     shown = 'ἠ'
			case Iota:    shown = 'ἰ'
			case Omicron: shown = 'ὀ'
			case Upsilon: shown = 'ὐ'
			case Omega:   shown = 'ὠ'
		}

		if g.breath == Rough {
			shown += 1
		}

		switch g.accent {
			case Acute:      shown += 4
			case Grave:      shown += 2
			case Circumflex: if low != Omicron { shown += 6 }
		}

		if g.iota {
			switch low {
				case Alpha: shown += 128
				case Eta:   shown += 112
				case Omega: shown += 64
			}
		}
	} else if g.iota {
		switch low {
			case Alpha: shown = 'ᾳ'
			case Eta:   shown = 'ῃ'
			case Omega: shown = 'ῳ'
		}

		switch g.accent {
			case Acute:      shown += 1
			case Grave:      shown -= 1
			case Circumflex: shown += 4
		}
	} else if g.diaeresis {
		var (
			grave rune
			tonos rune
		)

		if low == Iota {
			shown, grave, tonos = 'ϊ', 'ῒ', 'ΐ'
		} else if low == Upsilon {
			shown, grave, tonos = 'ϋ', 'ῢ', 'ΰ'
		}

		if g.accent != Unaccented {
			switch g.accent {
				case Acute:      shown = grave + 1
				case Grave:      shown = grave
				case Circumflex: shown = grave + 5
				case Tonos:      shown = tonos
			}
		}
	} else if g.accent != Unaccented {
		switch g.accent {
			case Circumflex:
				switch low {
					case Alpha:   shown = 'ᾶ'
					case Eta:     shown = 'ῆ'
					case Iota:    shown = 'ῖ'
					case Upsilon: shown = 'ῦ'
					case Omega:   shown = 'ῶ'
				}
			case Tonos:
				switch low {
					case Alpha:   shown = 'ά'
					case Epsilon: shown = 'έ'
					case Eta:     shown = 'ή'
					case Iota:    shown = 'ί'
					case Omicron: shown = 'ό'
					case Upsilon: shown = 'ύ'
					case Omega:   shown = 'ώ'
				}
			default:
				switch low {
					case Alpha:   shown = 'ὰ'
					case Epsilon: shown = 'ὲ'
					case Eta:     shown = 'ὴ'
					case Iota:    shown = 'ὶ'
					case Omicron: shown = 'ὸ'
					case Upsilon: shown = 'ὺ'
					case Omega:   shown = 'ὼ'
				}

			if g.accent == Acute {
				shown += 1
			}

		}
	} else if g.length != Assumed {
		switch low {
			case Alpha:   shown = 'ᾰ'
			case Iota:    shown = 'ῐ'
			case Upsilon: shown = 'ῠ'
		}

		if g.length == Long {
			shown += 1
		}
	} 

	return shown
}
