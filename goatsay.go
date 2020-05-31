package main

import (
	"fmt"
	"strings"
)

func makeBubble(text []string, longest int) string {
	var out []string
	lines := len(text)
	borders := []string{"/", "\\", "|", "(", ")"}
	top := " " + strings.Repeat("_", longest+2)
	bottom := " " + strings.Repeat("-", longest+2)

	out = append(out, top)

	if lines == 1 {
		s := fmt.Sprintf("%s %s %s", borders[3], text[0], borders[4])
		out = append(out, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], text[0], borders[1])
		out = append(out, s)
		i := 1
		for ; i < lines-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[2], text[i], borders[2])
			out = append(out, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[1], text[i], borders[0])
		out = append(out, s)
	}

	out = append(out, bottom)
	return strings.Join(out, "\n")
}
