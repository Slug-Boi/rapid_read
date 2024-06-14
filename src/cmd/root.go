/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rapid_read",
	Short: "A program that will allow you to read text files quickly and efficiently.",
	Long: `Rapid Read is a program that will allow you to read text files quickly and efficiently. 
	The idea behind the program is to allow the user to read text files at a faster pace than normal by
	using word grouping and text highlighting to help the user read the text faster. The GUI part is 
	customizable using the settings panel or the config file that is created when the program is first run.
	
	The GUI application will open by running the program with no arguments. If you want to open the program with
	a specific file you can run the program with the open command and the file path as an argument.
	
	The program also allows for PDF to text conversion using a specified tool in the config file. 
	The default tool is xPDFReaders pdftotext CLI tool.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 


	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


