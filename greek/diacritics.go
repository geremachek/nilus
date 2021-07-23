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
