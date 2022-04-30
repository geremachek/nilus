package main

import "strings"

// Parse the Beta Codes

func parse(raw string) string {
	var (
		chars = []rune(raw)
		end = len(chars)-1

		greek strings.Builder

		ch rune
	)

	for i := 0; i <= end; i++ {
		ch = chars[i]

		switch ch {

			// Convert symbols into diacritics

			case '/':  greek.WriteRune('\u0301') // Acute 
			case '\\': greek.WriteRune('\u0300') // Grave
			case '=':  greek.WriteRune('\u0342') // Circumflex
			
			case '(':  greek.WriteRune('\u0314') // Rough Breathing
			case ')':  greek.WriteRune('\u0313') // Smooth Breathing
		
			case '+':  greek.WriteRune('\u0308') // Diaeresis
			case '|':  greek.WriteRune('\u0345') // Iota Subscript

			case '&':  greek.WriteRune('\u0304') // Macron
			
			case '\'': greek.WriteRune('\u0306') // Breve
			
			case ':':  greek.WriteRune('\u00b7') // Colon
			case ';':  greek.WriteRune('\u037e') // Question Mark
			case '_':  greek.WriteRune('\u2014') // Dash
			case '#':  greek.WriteRune('\u02b9') // Keraia
		
			default:
				var gramma rune

				// Check to see if we have a multi-character Beta Code

				if i <= end-1 {
					for j := 3; j >= 1; j-- { // Loop backwards through a character cluster, searching for a valid code
						gramma = fromBetaCode(string(chars[i:i+j]))

						if gramma != 0 { // If we generate a valid gramma, exit the loop
							i += j-1 // Jump past the skipped characters
							break
						}
					}
				} else { // Otherwise parse characters individuallly
					gramma = fromBetaCode(string(ch))
				}

				if gramma != 0 {
					greek.WriteRune(gramma)
				} else {
					greek.WriteRune(ch)
				}
		}
	}

	return greek.String()
}
