package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kinpoko/ght/ghtrend"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ght",
	Short: "GitHub Trend",
	Long:  ``,

	RunE: func(cmd *cobra.Command, args []string) error {
		language, err := cmd.Flags().GetString("language")

		if err != nil {
			return err
		}

		res, err := ghtrend.Scraping(language)

		if err != nil {
			return err
		}

		for _, i := range res {
			color.Blue(i.Url)
			color.Green(i.Name)
			fmt.Println(i.Description)
			fmt.Println(i.Language)
			fmt.Print("\n")
		}

		return nil

	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("language", "l", "", "programming anguage")
}
