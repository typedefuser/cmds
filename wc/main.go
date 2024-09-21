package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "wc [file]",
		Short: "Word Count CLI",
		Long:  `A simple word count CLI tool similar to the Unix wc command.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			wordCount(args[0])
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func wordCount(input string) {
	var content []byte
	var filename string
	if _, err := os.Stat(input); err == nil {
		filename = input
		content, err = os.ReadFile(input)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			return
		}
	} else {
		content = []byte(input)
	}

	text := string(content)
	lines := strings.Split(text, "\n")
	words := strings.Fields(text)
	bytes := len(content)

	fmt.Printf("Lines:%d Words:%d Bytes:%d\nFilename:%s\n", len(lines), len(words), bytes, filename)
}
