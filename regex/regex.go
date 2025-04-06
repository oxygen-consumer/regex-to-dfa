package regex

import "fmt"

type Regex struct {
	Alphabet map[Symbol]bool
	Infix    []Token
}

func (regex *Regex) GeneratePostfix() {
	// TODO:
}

func (regex Regex) Print() {
	fmt.Println("Alphabet:")
	for sym := range regex.Alphabet {
		fmt.Printf(" - %c\n", sym)
	}

	fmt.Print("\nInfix:")
	for _, tok := range regex.Infix {
		fmt.Printf(" %c", tok.Value)
	}

	fmt.Println()
}
