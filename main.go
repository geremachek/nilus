package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"strings"
)

const betaTable string = `Greek Beta Code
---------------
    α A
    β B
    γ G
    δ D
    ε E
    ϝ V
    ζ Z
    η H
    Θ Q
    ι I
    κ K
    λ L
    μ M
    ν N
    ξ C
    ο O
    π P
    ρ R
    σ S, S1
    ς S, S2, J
    ϲ S3
    τ T
    υ U
    φ F
    χ X
    ψ Y
    ω W

*A -> Α

Puncuation Beta Code
--------------------
         . .
	 , ,
	 · :
	 ; ;
	 ‐ -
	 — _
	 ʹ #

        Diacritic BetaCode
	------------------
 Smooth Breathing )
  Rough Breathing (
    Accute Accent /
Circumflex Accent =
     Grave Accent \
        Diaeresis +
   Iota Subscript |
           Macron &
	    Breve '`

func main() {
	newline := flag.Bool("n", false, "Don't Print a newline when argumentss are passed")	
	table := flag.Bool("t", false, "Display a Beta Code usage table")

	flag.Parse()

	args := flag.Args()

	if *table {
		fmt.Println(betaTable)
	} else {
		if len(args) > 0 { // parse arguments...
			fmt.Print(parse(strings.Join(args, " ")))

			if !*newline {
				fmt.Print("\n")
			}
		} else { // parse stdin
			scanner := bufio.NewScanner(os.Stdin)

			for scanner.Scan() {
				fmt.Println(parse(scanner.Text()))
			}
		}
	}
} 
