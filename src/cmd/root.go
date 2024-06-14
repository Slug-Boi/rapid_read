/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
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

func Setup_config(input, confDir string) {
	// Create config directory and file
	if input == "y" {

		// Create config directory
		err := os.Mkdir(confDir+"rapid_read", 0755)
		if err != nil {
			panic("Failed to create config directory: " + err.Error())
		}

		// Create config file
		file, err := os.Create(confDir + "rapid_read/config.json")
		if err != nil {
			panic("Failed to create config file: " + err.Error())
		}

		// Write to config file
		_, err = file.WriteString("PDFTool: \"pdftotext\" \n")
		if err != nil {
			panic("Failed to write to config file: " + err.Error())
		}

		fmt.Println("Config file created, default PDF to text converter is pdftotext.")

	} else {
		fmt.Println("Config file not found, exiting...")
		os.Exit(1)
	}

	Continue()
}

func Check_config() *Config {
	confDir, err := os.UserConfigDir()
	if err != nil {
		panic("User $HOME variable not set: " + err.Error())
	}

	_, err = os.Stat(confDir + "rapid_read/config.json")
	if err != nil {
		fmt.Println("Config file not found, would you like the program to create one for you? (y/n)")
		var input string
		fmt.Scanln(&input)

		Setup_config(input, confDir)
	}

	return unmarshal_config(confDir)
}

func Continue() {
	fmt.Println("Would you like to continue? (y/n)")
	var input string
	fmt.Scanln(&input)
	if input != "y" {
		os.Exit(1)
	}
}

type Config struct {
	PDFTool        string
	Font           string
	FontSize       int
	HighlightColor string
	HighlightSpeed int
	WordAmount     int
	DarkMode       bool
	Keybinds       map[string]string
}

func unmarshal_config(confDir string) *Config {
	// Read config file
	var config Config
	file, err := os.ReadFile(confDir + "rapid_read/config.json")
	if err != nil {
		panic("Failed to open config file: " + err.Error())
	}
	if err := json.Unmarshal(file, &config); err != nil {
		panic("Failed to unmarshal config file: " + err.Error())
	}

	return &config
}
