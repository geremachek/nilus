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
	HookedUpsilon = 'ϒ'

	Heta          = 'ͱ'
	Sampi         = 'ͳ'
	Digamma       = 'ͷ'
	Antisigma     = 'ͻ'
	DottedSigma   = 'ͼ'
	DottedAnti    = 'ͽ'
	Jot           = 'ϳ'
	Kai           = 'ϗ'
	ABeta         = 'ϐ'
	ATheta        = 'ϑ'
	APhi          = 'ϕ'
	APi           = 'ϖ'
	Qoppa         = 'ϙ'
	Stigma        = 'ϛ'
	Waw           = 'ϝ'
	AQoppa        = 'ϟ'
	ASampi        = 'ϡ'
	AKappa        = 'ϰ'
	ARho          = 'ϱ'
	ASigma        = 'ϲ'
	ABTheta       = 'ϴ'
	AEpsilon      = 'ϵ'
	AREpsilon     = '϶'
	Sho           = 'ϸ'
	San           = 'ϻ'
	StrokeRho     = 'ϼ'
)

func fromLatin(latin rune, arch bool) rune {
	var (
		upper, letter = handleCase(latin)
		greek  = ' '
	)
	if arch {
		switch letter {
			case 'h': greek = Heta
			case 'i': greek = Sampi
			case 'w': greek = Digamma
			case 'c': greek = Antisigma
			case 'z': greek = DottedSigma
			case 'x': greek = DottedAnti
			case 'j': greek = Jot
			case 'm': greek = Kai
			case 'b': greek = ABeta
			case 'u': greek = ATheta
			case 'f': greek = APhi
			case 'p': greek = APi
			case 'q': greek = Qoppa
			case 't': greek = Stigma
			case 'v': greek = Waw
			case 'a': greek = AQoppa
			case 'o': greek = ASampi
			case 'k': greek = AKappa
			case 'r': greek = ARho
			case 's': greek = ASigma
			case 'y': greek = ABTheta
			case 'e': greek = AEpsilon
			case 'd': greek = AREpsilon
			case 'n': greek = Sho
			case 'g': greek = San
			case 'l': greek = StrokeRho
		}
	} else {
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
			case 'q': greek = HookedUpsilon
		}
	}

	if upper { greek = unicode.ToUpper(greek) }

	return greek
}

func handleCase(c rune) (bool, rune) {
	return unicode.IsUpper(c), unicode.ToLower(c)
}
