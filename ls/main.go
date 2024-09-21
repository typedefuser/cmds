package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ls",
		Short: "List files",
		Long:  `A simple ls command that lists files in the current directory.`,
		Args:  cobra.MaximumNArgs(1),
		Example: `ls
				   ls [directory]`,
		Run: func(cmd *cobra.Command, args []string) {
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			showfiles(dir)
		},
	}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}
func showfiles(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	for _, file := range files {
		fileinfo, err := file.Info()
		if err != nil {
			fmt.Print(err)
			panic(err)
		}
		fmt.Println("filename size")
		fmt.Printf("%s\t %d\n", file.Name(), fileinfo.Size())
	}
}
