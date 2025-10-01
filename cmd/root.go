package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "todo",
	Short: "Простое приложение для управления задачами",
	Long: "todo - это CLI-приложение для управления списком дел",
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}