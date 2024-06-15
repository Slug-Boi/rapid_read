package cmd

import (
	"fmt"
	"os/exec"
	"strings"

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

		// Get PDFTool from config file
		config := Check_config()

		PDFTool = config.PDFTool

		if PDFTool == "" {
			panic("PDFTool not set in config file")
		}

		pdfCmd := exec.Command(PDFTool, args[0])

		err := pdfCmd.Run()
		if err != nil {
			panic("Failed to convert PDF file: " + err.Error())
		}
		fmt.Println(PDFTool + " called")
	
		// Open the text file	
		nameLst := strings.Split(args[0], ".")
		name := nameLst[len(nameLst)-2]
		Open(name + ".txt")
	},
}



func init() {
	rootCmd.AddCommand(pdfCmd)

}
