package cdk

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var GenerateSubCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate source/destination template",
	Long:  `Generate source/destination`,
	Run: func(cmd *cobra.Command, args []string) {
		var answer;

		// Get the language to generate source
		var languageQuestion = []*survey.Question{
			{
				Name: "language",
				Prompt: &survey.Select{
					Message: "Choose language:",
					Options: Languages,
					Default: "python",
				},
			},
		}

		answer = struct {
			Language string
		}{}

		// perform the questions
		err := survey.Ask(languageQuestion, &answer)

		if err != nil {
			cobra.CheckErr(err)
			return
		}

		
	},
}
