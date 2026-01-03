package main

import (
	"fmt"
	"github-activity/cmd"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
