package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var quads = map[string][]string{
	"quadA": {"o-o", "| |", "o-o"},
	"quadB": {"/-\\", "| |", "\\-/"},
	"quadC": {"A", "B", "C"},
	"quadD": {"A", "B", "C"},
	"quadE": {"A", "B", "C"},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	var pattern_matches []string
	for name, pattern := range quads {
		if match(input, pattern) {
			pattern_matches = append(pattern_matches, fmt.Sprintf("[%s] [%d] [%d]", name, len(input[0]), len(input)))
		}
	}

	if len(pattern_matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		sort.Strings(pattern_matches)
		fmt.Println(strings.Join(matches, " || "))
	}
}

func match(input, pattern []string) bool {
	return strings.Join(input, "\n") == strings.Join(pattern, "\n")
}
