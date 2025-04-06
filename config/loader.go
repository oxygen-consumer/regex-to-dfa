package config

import (
	"bufio"
	"fmt"
	"os"
	"regex-to-dfa/regex"
	"strings"
	"unicode"
)

func LoadRegexFromConfig(filePath string) (*regex.Regex, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	re := &regex.Regex{
		Alphabet: make(map[regex.Symbol]bool),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, "#")[0]
		line = strings.TrimSpace(line)

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "alphabet":
			for s := range strings.SplitSeq(value, ",") {
				s = strings.TrimSpace(s)

				if len(s) != 1 {
					return nil, fmt.Errorf("invalid symbol: '%s'", s)
				}

				re.Alphabet[regex.Symbol(s[0])] = true
			}

		case "regex":
			runes := []rune(value)
			for _, r := range runes {
				if unicode.IsSpace(r) {
					continue
				}
				if regex.IsOperator(r) || r == '(' || r == ')' {
					continue
				}
				if _, exists := re.Alphabet[regex.Symbol(r)]; !exists {
					return nil, fmt.Errorf("symbol not in alphabet: %c", r)
				}
			}

			re.Infix, err = regex.Tokenize(runes, re.Alphabet)
			if err != nil {
				return nil, err
			}

		default:
			fmt.Printf("WARN: unknown config key: %s\n", key)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return re, nil
}
