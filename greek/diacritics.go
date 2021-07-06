package greek

type Accent rune

const (
	Accute Accent = '''
	Grave         = '`'
	Circumflex    = '^'
)

type Breathing int

const (
	Rough Breathing = iota
	Smooth
	None
)
