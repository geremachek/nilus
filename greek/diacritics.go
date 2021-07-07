package greek

type Accent int

const (
	Acute Accent = iota
	Grave
	Circumflex
	Tonos
	Unaccented
)

type Breathing int

const (
	Rough Breathing = iota
	Smooth
	Unmarked
)

type Length int

const (
	Long Length = iota
	Short
	Assumed
)
