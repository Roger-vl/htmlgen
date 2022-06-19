package cmd

import (
	"github.com/RogerDurdn/htmlgen/core"
	"os"

	"github.com/spf13/cobra"
)

var (
	templatePath, dataPath, fileIdentifier string
	size                                   int8
)

var rootCmd = &cobra.Command{
	Use:   "htmlgen",
	Short: "Html generator by templates",
	Long: `
Html generator by templates:

You have to provide a template and the data that has to be used on the 
generation of templates, this files are html(template) and json(data)

-- example: htmlgen -t myTemplate.html -d myData.json 

-by: Roger VL`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Generate(templatePath, dataPath, fileIdentifier, size)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&templatePath, "template", "t", "", "file Name of the html template to be used (required)")
	rootCmd.Flags().StringVarP(&dataPath, "data", "d", "", "file Name of the json data to be used (required)")
	rootCmd.Flags().StringVarP(&fileIdentifier, "fileIdentifier", "i", "template-", "identifier on data to use like name on the generated file")
	rootCmd.Flags().Int8VarP(&size, "size", "s", 0, "size of the generated templates")
	rootCmd.MarkFlagRequired("template")
	rootCmd.MarkFlagRequired("data")
}
