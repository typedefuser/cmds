/*
 * Project: cat
 * File: main.go
 * Description: A simple linux command utility to read,append files
 * Author: Nikhilesh Majhi
 *
 * Usage:
 * - Compile: go build
 * - Run: ./cat
 *
 * Example:
 * - cat file1.txt
 * - cat file1.txt file2.txt > output.txt
 */
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cat [file]",
		Short: "CAT cli",
		Long:  "A simple linux command utility to read,append files",
		Args:  cobra.MinimumNArgs(1),
		Run:   runcat,
	}
	_ = rootCmd.Execute()
}
func runcat(cmd *cobra.Command, args []string) {
	switch len(args) {
	case 1:
		fmt.Print(string(readFile(args[0])))
	case 2:
		appendFiles(args[0], args[1])
	default:
		cmd.Help()
	}
}

func readFile(filename string) []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return content
}

func appendFiles(filename string, file2 string) {
	str1 := readFile(filename)
	str2 := readFile(file2)
	fmt.Print(string(str1))
	fmt.Println("")
	fmt.Print(string(str2))
}
