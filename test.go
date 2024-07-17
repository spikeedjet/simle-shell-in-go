"use go write shell dont do any low-level stuff as managing the memory ourselves
its use the package os and os exec do the dirty work"


package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// execInput executes the command and its arguments
func execInput(input string) error {
	// Remove leading and trailing whitespace
	input = strings.TrimSpace(input)

	// Split the input into command and arguments
	args := strings.Split(input, " ")
    fmt.Println(args)
    fmt.Println(strings.Join(args, " "))


	// For other commands, execute them directly
    cmd := exec.Command("cmd", "/C", strings.Join(args, " "))
	// cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command
	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		// Read the keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		// Handle the execution of the input
		if err := execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, "Error executing command:", err)
		}
	}
}
