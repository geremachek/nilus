package main

import (
	"unicode"
	"strings"
)

// Beta Code Table

var betaCodes = map[string]rune {
	"A":  'α',
	"B":  'β',
	"G":  'γ',
	"D":  'δ',
	"E":  'ε',
	"V":  'ϝ',
	"Z":  'ζ',
	"H":  'η',
	"Q":  'θ',
	"I":  'ι',
	"K":  'κ',
	"L":  'λ',
	"M":  'μ',
	"N":  'ν',
	"C":  'ϲ',
	"O":  'ο',
	"P":  'π',
	"R":  'ρ',
	"S":  'σ',
	"S1": 'σ',
	"S2": 'ς',
	"J":  'ς',
	"S3": 'ϲ',
	"T":  'τ',
	"U":  'υ',
	"F":  'φ',
	"X":  'χ',
	"Y":  'ψ',
	"W":  'ω',
}

// Non-alphabetic Greek symbols

var nonAlphabetic = map[rune]rune {
	'/':  '\u0301', // Acute
	'\\': '\u0300', // Grave
	'=':  '\u0342', // Circumflex

	'(':  '\u0314', // Rough Breathing
	')':  '\u0313', // Smooth Breathing

	'+':  '\u0308', // Diaeresis
	'|':  '\u0345', // Iota Subscript

	'&':  '\u0304', // Macron
	'\'': '\u0306', // Breve

	':':  '\u00b7', // Colon
	';':  '\u037e', // Question Mark
	'_':  '\u2014', // Dash
	'#':  '\u02b9', // Keraia
}

// Convert a Beta Code to a letter of the Greek Alphabet

func fromBetaCode(bc string) (gramma rune) {
	var (
		key string = strings.ToUpper(bc)
		runes []rune = []rune(key)
	)

	if len(runes) > 1 && runes[0] == '*' {
		key = string(runes[1:])
	}

	// If we see an asterisk, the letter will be capitalized

	if value, validKey := betaCodes[key]; validKey {
		gramma = value

		if runes[0] == '*' {
			gramma = unicode.ToUpper(value)
		}
	}

	return
}
