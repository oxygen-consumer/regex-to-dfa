package main

import (
	"flag"
	"fmt"
	"os"
	"regex-to-dfa/config"
)

func main() {
	var filePath string

	flag.StringVar(&filePath, "config", "", "Path to the config file")
	flag.Parse()

	if filePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	regex, err := config.LoadRegexFromConfig(filePath)
	if err != nil {
		fmt.Printf("Error parsing config file: %s\n", err.Error())
		os.Exit(1)
	}

	regex.Print()
}
