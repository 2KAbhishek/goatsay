package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
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

func normalizeStr(text []string, longest int) []string {
	var out []string
	for _, line := range text {
		s := line + strings.Repeat(" ", longest-utf8.RuneCountInString(line))
		out = append(out, s)
	}
	return out
}

func printArt(name string) {

	var goat = `             / /
          (\/_//')
           /   '/
          0  0   \
         /        \
        /    __/   \
       /,  _/ \     \_
       '-./ )  |     ~^~^~^~^~^~^~^~\~.
           (   /                     \_}
              |               /      |
              ;     |         \      /
               \/ ,/           \    |
               / /~~|~|~~~~~~|~|\   |
              / /   | |      | | '\ \
             / /    | |      | |   \ \
            / (     | |      | |    \ \
           /,_)    /__)     /__)   /,_/`

	var dog = `                               _
                            ,:'/   _..._
                           // ( '""-.._.'
                           \| /    6\___
                           |     6      4
                           |            /
                           \_       .--'
                           (_'---'')
                           / ''---'()
                         ,'        |
         ,            .''          |
         )\       _.-'             ;
        / |    .''   _            /
      /' /   .'       '.        , |
     /  /   /           \   ;   | |
     |  \  |            |  .|   | |
      \  '"|           /.-' |   | |
       '-..-\       _.;.._  |   |.;-.
             \    <'.._  )) |  .;-. ))
             (__.  '  ))-'  \_    ))'
                 ''--"'       '"""'`

	switch name {
	case "goat":
		fmt.Println(goat)
	case "dog":
		fmt.Println(dog)
	default:
		fmt.Println("Art not found.")
	}
}

func main() {
	var figure string
	flag.StringVar(&figure, "f", "goat", "Figure name - options goat, dog")
	flag.Parse()

	inputInfo, _ := os.Stdin.Stat()
	if inputInfo.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("goatsay works with pipe inputs.\nUsage: command | goatsay")
		os.Exit(1)
	}

	var text []string

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		text = append(text, string(line))
	}

	text = tabsToSpaces(text)
	longest := getLongest(text)
	text = normalizeStr(text, longest)
	bubble := makeBubble(text, longest)

	fmt.Println(bubble)
	printArt(figure)
	fmt.Println()
}
