package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"push-swap/internal"
)

// main is the entry point of the push-swap checker.
// It delegates all logic to run, passing the real OS streams.
func main() {
	os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}

// run contains the core logic of the checker, accepting injectable streams
// so it can be tested without spawning a subprocess.
// Returns 0 on success, 1 on any error.
func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	// Parse and validate the initial integer values from CLI arguments.
	values, err := internal.ParseArgs(args)
	if err != nil {
		fmt.Fprintln(stderr, "Error")
		return 1
	}

	// Nothing to sort if no values were provided.
	if len(values) == 0 {
		return 0
	}

	// Initialise the two-stack state with the parsed values loaded into stack A.
	state := internal.NewState(values)

	// Read and execute instructions one per line from stdin.
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if err := internal.ExecuteInstruction(state, line); err != nil {
			fmt.Fprintln(stderr, "Error")
			return 1
		}
	}

	// Check for any scanner-level I/O error after the loop.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(stderr, "Error")
		return 1
	}

	// The solution is correct only when stack A is fully sorted and stack B is empty.
	if internal.IsSortedAscending(state.A.Values()) && state.B.IsEmpty() {
		fmt.Fprintln(stdout, "OK")
		return 0
	}

	fmt.Fprintln(stdout, "KO")
	return 0
}
