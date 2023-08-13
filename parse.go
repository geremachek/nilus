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

		// add the non-alphabetic character to the builder if it is valid.

		if na, v := nonAlphabetic[ch]; v {
			greek.WriteRune(na)
		} else {
			var gramma rune

			// Check to see if we have a multi-character Beta Code

			if i <= end-1 {
				for j := end-i+1; j >= 1; j-- { // Loop backwards through a character cluster, searching for a valid code
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
