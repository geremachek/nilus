package main

// accent variations

type Accent int

const (
	Unaccented Accent = iota
	Acute
	Grave
	Circumflex
	Tonos
)

// breath variations

type Breathing int

const (
	
	Unmarked Breathing = iota
	Rough
	Smooth
)

// length variations

type Length int

const (
	Assumed Length = iota
	Long
	Short
)

// make sure we can use this kind of diacritic...

func validCircumflex(ch rune) bool {
	return ch == Alpha || ch == Upsilon || ch == Iota || ch == Omega || ch == Eta
}

func validIota(ch rune) bool {
	return ch == Alpha || ch == Eta || ch == Omega
}

func validDiaeresis(ch rune) bool {
	return ch == Iota || ch == Upsilon
}
