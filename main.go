package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"strings"
)

func main() {
	newline := flag.Bool("n", false, "Don't Print a newline when argumentss are passed")	

	flag.Parse()

	args := flag.Args()

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
