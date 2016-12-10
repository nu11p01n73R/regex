package main

import (
	"fmt"
)

// Parses the pattern string.
// The input pattern is parsed into multiple smaller
// patterns which can be any of
// 		- simple, eg "abcd"
// 		- star closure, eg "a*"
// 		- plus closure, eg "a+"
// Return
// 		[]string Array of string corresponding to the patterns.
func parsePattern(s string) []string {
	var patterns []string
	var pattern string
Loop:
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			switch s[i+1] {
			case '+', '*':
				if len(pattern) > 0 {
					patterns = append(patterns, pattern)
				}
				pattern = ""
			}
		}
		pattern += string(s[i])

		// Patterns ending with *|+
		switch s[i] {
		case '+', '*':
			patterns = append(patterns, pattern)
			pattern = ""
			continue Loop
		}
	}
	// Append the last pattern
	if len(pattern) > 0 {
		patterns = append(patterns, pattern)
	}
	return patterns
}

// Matches * patterns
// Example
// 		a*
// Return
// 		bool 	if it can be matched from the current start.
// 		int 	position where the pattern doesn't match.
func starClosure(s string, p string, start int) (bool, int) {
	char := p[0]
	i := start
	for ; i < len(s) && (s[i] == char || char == '.'); i = i + 1 {
	}

	return true, i
}

// Matches + patterns
// Example
// 		a+
// Return
// 		bool 	if it can be matched from the current start.
// 		int 	position where the pattern doesn't match.
func plusClosure(s string, p string, start int) (bool, int) {
	char := p[0]
	i := start
	for j := 0; i < len(s) && j < len(p) && (s[i] == char || char == '.'); i, j = i+1, j+1 {
	}
	if i > start {
		return true, i
	} else {
		return false, start
	}
}

// Does a string match.
// This is for matches with single literals,
// for example
// 		"abcd"
// type of patterns
// Return
// 		bool 	if it can be matched from the current start.
// 		int 	position where the pattern doesn't match.
func patternMatch(s string, p string, start int) (bool, int) {
	i := start
	j := 0
	for ; i < len(s) && j < len(p) && (s[i] == p[j] || p[j] == '.'); i, j = i+1, j+1 {
	}

	if j < len(p) {
		return false, start
	} else {
		return true, i
	}
}

// Matches a pattern to the string.
// return boolean, wether the pattern matches the
// string or not.
func isMatch(s string, p string) bool {
	patterns := parsePattern(p)
	ptr := 0
	match := false

	for _, pattern := range patterns {
		if len(pattern) == 2 && pattern[1] == '*' {
			match, ptr = starClosure(s, pattern, ptr)
		} else if len(pattern) == 2 && pattern[1] == '+' {
			match, ptr = plusClosure(s, pattern, ptr)
		} else {
			match, ptr = patternMatch(s, pattern, ptr)
		}

		if !match {
			return false
		}
	}
	return match && ptr == len(s)
}

func main() {
	fmt.Println(isMatch("aaaaaabcd", "a*bcde"))
}
