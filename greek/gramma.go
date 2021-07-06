package greek

type Gramma struct {
	letter rune
	accent Accent

	breath Breathing
	long bool
	diaeresis bool
	iota bool
}

func NewGramma(let rune, a accent, b Breathing, l, d, i bool) {
	return Gramma { let, a, b, l, d, i }
}

func (g Gramma) Show() (shown rune) {
	upper, low := handleCase(g.letter)
	shown = low

	if g.breath != None {
		switch low {
			case Alpha;   shown = 'ἀ'
			case Epsilon: shown = 'ἐ'
			case Eta:     shown = 'ἠ'
			case Iota:    shown = 'ἰ'
			case Omicron: shown = 'ὀ'
			case Upsilon: shown = 'ὐ'
			case Omega:   shown = 'ὠ'
			case Rho:     shown = 'ῤ'
		}

		if g.breath == Rough {
			shown += 1
		}

		switch g.accent {
			case Acute:      shown += 4
			case Grave:      shown += 2
			case Circumflex: shown += 6
		}
	}

	return shown
}
