package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	gr "github.com/geremachek/nilus/greek"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(gr.Parse(strings.Join(os.Args[1:], " ")))
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			fmt.Println(gr.Parse(scanner.Text()))
		}
	}
}
