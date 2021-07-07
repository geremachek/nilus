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
	arch := flag.Bool("a", false, "Use alternate/archaic characters")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		fmt.Println(gr.Parse(strings.Join(args, " "), *arch))
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			fmt.Println(gr.Parse(scanner.Text(), *arch))
		}
	}
}
