package greek

// accent variations

type Accent int

const (
	Acute Accent = iota
	Grave
	Circumflex
	Tonos
	Unaccented
)

// breath variations

type Breathing int

const (
	Rough Breathing = iota
	Smooth
	Unmarked
)

// length variations

type Length int

const (
	Long Length = iota
	Short
	Assumed
)

// make sure we can use this kind of diacritic...

func validCircumflex(ch rune) bool {
	return ch == Alpha || ch == Upsilon || ch == Iota || ch == Omega
}

func validIota(ch rune) bool {
	return ch == Alpha || ch == Eta || ch == Omega
}

func validDiaeresis(ch rune) bool {
	return ch == Iota || ch == Upsilon
}
