package greek

import (
	"unicode"
	"strings"
)

// greek characters

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

// the common transliterations between both modes

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
	}

	return ' '
}

// convert latin to greek (unaccented)

func fromLatin(latin string, keyb bool) rune {
	var (
		chars = []rune(latin)
		checkCap = 0
		
		upper bool
		letter rune
		
		greek = ' '
		l = len(chars)
	)

	if chars[0] == '\\' && l == 2 {
		checkCap = 1
	}

	upper, letter = handleCase(chars[checkCap])

	if l < 2 {
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
				case 'w': greek = UltimateSigma
				case 'q': return  HookedUpsilon
			}
		} else {
			switch strings.ToLower(latin) {
				case "\\e": greek = Eta
				case "th":  greek = Theta
				case "x":   greek = Xi
				case "u":   greek = Upsilon
				case "y":   greek = Upsilon
				case "\\y": return  HookedUpsilon
				case "ch":  greek = Chi
				case "ps":  greek = Psi
				case "\\o": greek = Omega
				case "ph":  greek = Phi
				case "\\s": greek = UltimateSigma
				case "rh":  greek = 'ῥ'
			}
		}
	}
	
	if upper { greek = unicode.ToUpper(greek) }

	return greek
}

// lower the character, return true if the character is uppercase

func handleCase(c rune) (bool, rune) {
	return unicode.IsUpper(c), unicode.ToLower(c)
}
