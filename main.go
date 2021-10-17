package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"strings"
)

func main() {
	key := flag.Bool("k", false, "Use keyboard substitution mode")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 { // parse arguments...
		fmt.Println(parse(strings.Join(args, " "), *key))
	} else { // parse stdin
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			fmt.Println(parse(scanner.Text(), *key))
		}
	}
}
