package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sourcegraph/conc/panics"
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:   "pdf",
	Short: "Convert and open a PDF file to read",
	Long: `The PDF command allows you to convert a PDF file to a text file. using (...) after
	the conversion the program will automatically read the text file and display it on the screen. 
	The PDF file will need to be OCR compatible. more than likely for the program to work.
	
	You can specify which PDF converting tool to use by editing the config file. The default is (...)`,
	Run: func(cmd *cobra.Command, args []string) {
		var PDFTool string

		confDir, err := os.UserConfigDir()
		if err != nil {
			panic("User $HOME variable not set: " + err.Error())
		}

		_, err = os.Stat(confDir + "rapid_read/config.json")
		if err != nil {
			fmt.Println("Config file not found, would you like the program to create one for you? (y/n)")
			var input string
			fmt.Scanln(&input)

			//TODO: Turn this into separate functions for readability and move it into root command file
			if input == "y" {
				err = os.Mkdir(confDir+"rapid_read", 0755)
				if err != nil {
					panic("Failed to create config directory: " + err.Error())
				}

				file , err := os.Create(confDir + "rapid_read/config.json")
				if err != nil {
					panic("Failed to create config file: " + err.Error())
				}

				_, err = file.WriteString("PDFTool: \"pdftotext\" \n")
				if err != nil {
					panic("Failed to write to config file: " + err.Error())
				}

				fmt.Println("Config file created, please edit the file to specify the PDF converting tool if you want to use a different tool than pdftotext.")

				fmt.Println("Would you like to continue with the default PDF converting tool? (y/n)")
				fmt.Scanln(&input)
				if input == "y" {
					PDFTool = "pdftotext"
				} else {
					fmt.Println("Exiting...")
					os.Exit(1)
				}
			} else {
				fmt.Println("Config file not found, exiting...")
				os.Exit(1)
			}
		}



		pdfCmd := exec.Command(PDFTool, args[0])

		err = pdfCmd.Run()
		if err != nil {
			panic("Failed to convert PDF file: " + err.Error())
		}
		fmt.Println(PDFTool + " called")


		
	},
}

func init() {
	rootCmd.AddCommand(pdfCmd)

}
