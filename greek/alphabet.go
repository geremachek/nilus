package greek

import "unicode"

const (
	Alpha         = 'α'
	Beta          = 'β'
	Gamma         = 'γ'
	Delta         = 'δ'
	Epsilon       = 'ε'
	Zeta          = 'ζ'
	Eta           = 'η'
	Theta         = 'θ'
	Iota          = 'ι'
	Kappa         = 'κ'
	Lambda        = 'λ'
	Mu            = 'μ'
	Nu            = 'ν'
	Xi            = 'ξ'
	Omicron       = 'ο'
	Pi            = 'π'
	Rho           = 'ρ'
	Sigma         = 'σ'
	Tau           = 'τ'
	Upsilon       = 'υ'
	Phi           = 'φ'
	Chi           = 'χ'
	Psi           = 'ψ'
	Omega         = 'ω'
	UltimateSigma = 'ς'

)

func fromLatin(latin rune) rune {
	var (
		upper, letter = handleCase(latin)
		greek  = ' '
	)

	switch letter {
		case 'a': greek = Alpha
		case 'b': greek = Beta
		case 'g': greek = Gamma
		case 'd': greek = Delta
		case 'e': greek = Epsilon
		case 'z': greek = Zeta
		case 'h': greek = Eta
		case 'u': greek = Theta
		case 'i': greek = Iota
		case 'k': greek = Kappa
		case 'l': greek = Lambda
		case 'm': greek = Mu
		case 'n': greek = Nu
		case 'j': greek = Xi
		case 'o': greek = Omicron
		case 'p': greek = Pi
		case 'r': greek = Rho
		case 's': greek = Sigma
		case 't': greek = Tau
		case 'y': greek = Upsilon
		case 'f': greek = Phi
		case 'x': greek = Chi
		case 'c': greek = Psi
		case 'v': greek = Omega
		case 'w': greek = UltimateSigma
	}

	if upper { greek = unicode.ToUpper(greek) }

	return greek
}

func handleCase(c rune) (bool, rune) {
	return unicode.IsUpper(c), unicode.ToLower(c)
}
