package main

import (
	"bytes"
	"strings"
	"testing"
)

// runTest is a helper that calls run with string-based stdin and captures stdout/stderr.
func runTest(args []string, stdin string) (stdout, stderr string, code int) {
	var outBuf, errBuf bytes.Buffer
	code = run(args, strings.NewReader(stdin), &outBuf, &errBuf)
	return strings.TrimSpace(outBuf.String()), strings.TrimSpace(errBuf.String()), code
}

// TestRun_OK verifies that a correctly sorted sequence prints "OK".
func TestRun_OK(t *testing.T) {
	// Input: [2 1] → "sa" swaps to [1 2] → sorted, B empty → OK.
	stdout, _, code := runTest([]string{"2", "1"}, "sa\n")
	if stdout != "OK" || code != 0 {
		t.Fatalf("expected OK/0, got %q/%d", stdout, code)
	}
}

// TestRun_KO verifies that an unsorted result prints "KO".
func TestRun_KO(t *testing.T) {
	stdout, _, code := runTest([]string{"3", "2", "1"}, "")
	if stdout != "KO" || code != 0 {
		t.Fatalf("expected KO/0, got %q/%d", stdout, code)
	}
}

// TestRun_NoArgs verifies that providing no arguments produces no output.
func TestRun_NoArgs(t *testing.T) {
	stdout, stderr, code := runTest([]string{}, "")
	if stdout != "" || stderr != "" || code != 0 {
		t.Fatalf("expected no output, got stdout=%q stderr=%q code=%d", stdout, stderr, code)
	}
}

// TestRun_InvalidArg verifies that a non-integer argument prints "Error" to stderr.
func TestRun_InvalidArg(t *testing.T) {
	_, stderr, code := runTest([]string{"abc"}, "")
	if stderr != "Error" || code != 1 {
		t.Fatalf("expected Error/1, got %q/%d", stderr, code)
	}
}

// TestRun_DuplicateArg verifies that duplicate values print "Error" to stderr.
func TestRun_DuplicateArg(t *testing.T) {
	_, stderr, code := runTest([]string{"1", "1"}, "")
	if stderr != "Error" || code != 1 {
		t.Fatalf("expected Error/1, got %q/%d", stderr, code)
	}
}

// TestRun_InvalidInstruction verifies that an unknown instruction prints "Error" to stderr.
func TestRun_InvalidInstruction(t *testing.T) {
	_, stderr, code := runTest([]string{"1", "2", "3"}, "invalid\n")
	if stderr != "Error" || code != 1 {
		t.Fatalf("expected Error/1, got %q/%d", stderr, code)
	}
}

// TestRun_AlreadySorted verifies that a pre-sorted sequence with no instructions prints "OK".
func TestRun_AlreadySorted(t *testing.T) {
	stdout, _, code := runTest([]string{"1", "2", "3"}, "")
	if stdout != "OK" || code != 0 {
		t.Fatalf("expected OK/0, got %q/%d", stdout, code)
	}
}
