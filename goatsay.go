package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func makeBubble(text []string, longest int) string {
	var out []string
	lineCount := len(text)
	borders := []string{"/", "\\", "|", "(", ")"}
	top := " " + strings.Repeat("_", longest+2)
	bottom := " " + strings.Repeat("-", longest+2)

	out = append(out, top)

	if lineCount == 1 {
		s := fmt.Sprintf("%s %s %s", borders[3], text[0], borders[4])
		out = append(out, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], text[0], borders[1])
		out = append(out, s)
		i := 1
		for ; i < lineCount-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[2], text[i], borders[2])
			out = append(out, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[1], text[i], borders[0])
		out = append(out, s)
	}

	out = append(out, bottom)
	return strings.Join(out, "\n")
}

func tabsToSpaces(text []string) []string {
	var out []string
	for _, line := range text {
		line = strings.ReplaceAll(line, "\t", "    ")
		out = append(out, line)
	}
	return out
}

func getLongest(text []string) int {
	longest := 0
	for _, line := range text {
		currLen := utf8.RuneCountInString(line)
		if currLen > longest {
			longest = currLen
		}
	}
	return longest
}
