package main

import "fmt"

/**
	https://www.codingame.com/ide/puzzle/format-string-validation
 **/
const (
	tilda = "~"
	mark  = "?"
)

func main() {
	text := "aaaaaauo"
	format := "~u?"

	if format == "" {
		fmt.Println("FAIL")
		return
	}

	if format == tilda {
		fmt.Println("MATCH")
		return
	}

	if format == mark && len(text) > 1 {
		fmt.Println("FAIL")
		return
	}

	tokens := make([]string, 0, len(format))
	for _, s := range format {
		tokens = append(tokens, string(s))
	}

	currentTokenPos := 0
	matched := true
	for _, r := range text {
		c := string(r)

		if currentTokenPos >= len(tokens)-1 {
			break
		}

		if tokens[currentTokenPos] == tilda {
			if tokens[currentTokenPos+1] == mark || tokens[currentTokenPos+1] == c {
				currentTokenPos++
			} else {
				continue
			}
		}

		if tokens[currentTokenPos] == mark || tokens[currentTokenPos] == c {
			currentTokenPos++
		} else {
			matched = false
		}
	}

	if currentTokenPos < len(tokens)-1 {
		for i := currentTokenPos + 1; i < len(tokens); i++ {
			if tokens[i] != tilda {
				matched = false
			}
		}
	}

	if matched {
		fmt.Println("MATCH")
	} else {
		fmt.Println("FAIL")
	}
}
