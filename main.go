package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"strings"
	gr "github.com/geremachek/nilus/greek"
)

func main() {
	key := flag.Bool("k", false, "Use keyboard substitution mode")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		fmt.Println(gr.Parse(strings.Join(args, " "), *key))
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			fmt.Println(gr.Parse(scanner.Text(), *key))
		}
	}
}
