package greek

import (
	"unicode"
	"strings"
)

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
	HookedUpsilon = 'ϒ'
)

func singleChar(ch rune) rune {
	switch ch {
		case 'a': return Alpha
		case 'b': return Beta
		case 'g': return Gamma
		case 'd': return Delta
		case 'e': return Epsilon
		case 'z': return Zeta
		case 'i': return Iota
		case 'k': return Kappa
		case 'l': return Lambda
		case 'm': return Mu
		case 'n': return Nu
		case 'o': return Omicron
		case 'p': return Pi
		case 'r': return Rho
		case 's': return Sigma
		case 't': return Tau
		case 'w': return UltimateSigma
	}

	return ' '
}

func fromLatin(latin string, keyb bool) rune {
	var (
		chars = []rune(latin)
		checkCap = 0
		
		upper bool
		letter rune
		
		greek = ' '
	)

	if chars[0] == '\\' {
		checkCap = 1
	}

	upper, letter = handleCase(chars[checkCap])

	if len(latin) < 2 {
		greek = singleChar(letter)
	}
	
	if greek == ' ' {
		if keyb {
			switch letter {
				case 'h': greek = Eta
				case 'u': greek = Theta
				case 'j': greek = Xi
				case 'y': greek = Upsilon
				case 'x': greek = Chi
				case 'c': greek = Psi
				case 'v': greek = Omega
				case 'f': greek = Phi
				case 'q': greek = HookedUpsilon
			}
		} else {
			switch strings.ToLower(latin) {
				case "\\e": greek = Eta
				case "th":  greek = Theta
				case "x":   greek = Xi
				case "u":   greek = Upsilon
				case "ch":  greek = Chi
				case "ps":  greek = Psi
				case "\\o": greek = Omega
				case "ph":  greek = Phi
				case "y":   greek = HookedUpsilon
			}
		}
	}
	
	if upper { greek = unicode.ToUpper(greek) }

	return greek
}

func handleCase(c rune) (bool, rune) {
	return unicode.IsUpper(c), unicode.ToLower(c)
}
