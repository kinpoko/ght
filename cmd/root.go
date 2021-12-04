package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/kinpoko/ght/ghtrend"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ght [programming language]",
	Short: "GitHub Trend",
	Long:  ``,

	Args: cobra.MaximumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		language := strings.Join(args, "")
		res, err := ghtrend.Scraping(language)

		if err != nil {
			return err
		}

		for _, i := range res {
			color.HiCyan(i.Url)
			color.HiGreen(i.Name)
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
}
