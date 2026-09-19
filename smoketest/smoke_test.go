// Package smoketest is this repository's real test coverage.
//
// Every solved problem under the repository root is its own `package main`
// with the computation written inline in main(), so there is nothing to
// import and unit-test directly. Instead these tests build and run two of
// the solutions as black-box programs -- D007 and D170 -- and feed them the
// exact sample input paiza.jp's own judge already scored 100 on (see
// D007/docment.md and D170/document.md), asserting the printed output
// matches the judge's own sample output. If a future edit to either
// solution changes its behavior, this test fails.
package smoketest

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// repoRoot returns the repository root, computed from this file's own path
// rather than the process's working directory, so `go test ./...` behaves
// the same regardless of where it is invoked from.
func repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("smoketest: could not determine this file's own path")
	}
	return filepath.Dir(filepath.Dir(file))
}

// runSolution builds and runs the `package main` in `dir` (relative to the
// repository root), feeding it `stdin` and returning what it printed.
func runSolution(t *testing.T, dir, stdin string) string {
	t.Helper()
	root := repoRoot(t)

	cmd := exec.Command("go", "run", "./"+dir)
	cmd.Dir = root
	cmd.Stdin = strings.NewReader(stdin)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf("go run ./%s (stdin %q) failed: %v\nstderr: %s", dir, stdin, err, stderr.String())
	}
	return stdout.String()
}

// TestD007PrintsNAsterisks exercises D007's own worked examples: for input
// N it must print exactly N '*' characters followed by a newline
// (D007/docment.md, input/output examples 1-3, scored 100 by paiza's judge).
func TestD007PrintsNAsterisks(t *testing.T) {
	cases := []struct {
		stdin string
		want  string
	}{
		{"4\n", "****\n"},
		{"6\n", "******\n"},
		{"2\n", "**\n"},
	}

	for _, tc := range cases {
		got := runSolution(t, "D007", tc.stdin)
		if got != tc.want {
			t.Fatalf("D007 with stdin %q printed %q, want %q", tc.stdin, got, tc.want)
		}
	}
}

// TestD170MultipliesLapDistanceByLapCount exercises D170's own worked
// examples: the total distance run is the lap distance N times the lap
// count M (D170/document.md, input/output examples 1-2, scored 100 by
// paiza's judge).
func TestD170MultipliesLapDistanceByLapCount(t *testing.T) {
	cases := []struct {
		stdin string
		want  string
	}{
		{"40\n15\n", "600\n"},
		{"125\n37\n", "4625\n"},
	}

	for _, tc := range cases {
		got := runSolution(t, "D170", tc.stdin)
		if got != tc.want {
			t.Fatalf("D170 with stdin %q printed %q, want %q", tc.stdin, got, tc.want)
		}
	}
}
