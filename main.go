package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/NickTaporuk/2021ai_test/parser"
	"golang.org/x/crypto/ssh/terminal"
)

// Stores the state of the terminal before making it raw
var regularState *terminal.State

// STEPS:
// 1. Validate
// 2. Parse args
// 3. Simultaneously create ast tree during process of parsing
// 4. computing tree go thru end's left and right branches simultaneously
func main() {

	if len(os.Args) > 1 {
		runAsOneTimeMode()
	} else {
		runAsInteractiveMode()
	}
}

func handleKey(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	if key == '\x03' {
		// Quit without error on Ctrl^C
		exit()
	}
	return "", 0, false
}

func exit() {
	// nolint
	terminal.Restore(0, regularState)
	fmt.Println()
	os.Exit(0)
}

func runAsOneTimeMode() {
	input := strings.Join(os.Args[1:], " ")
	r := strings.NewReader(input)
	p := parser.NewParser(r)

	state, err := p.Parse()
	if err != nil {
		panic(err)
	}

	data := state.String()
	fmt.Println(data)
}

func runAsInteractiveMode() {
	var err error
	regularState, err = terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	// nolint
	defer terminal.Restore(0, regularState)

	term := terminal.NewTerminal(os.Stdin, "scalc> ")
	term.AutoCompleteCallback = handleKey
	for {
		input, err := term.ReadLine()
		if err != nil {
			if err == io.EOF {
				// Quit without error on Ctrl^D
				exit()
			}
			panic(err)
		}

		if input == "exit" || input == "quit" || input == "q" {
			break
		}

		r := strings.NewReader(input)
		p := parser.NewParser(r)

		res, err := p.Parse()
		if err != nil {
			//nolint
			term.Write([]byte(fmt.Sprintln("Error: " + err.Error())))
			continue
		}

		// nolint
		term.Write([]byte(fmt.Sprintln(res)))
	}

}
