/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/kudagonbe/jpcal"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get day information.",
	Long:  "Get day information",
	Run: func(cmd *cobra.Command, args []string) {
		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			log.Fatal(err.Error())
		}

		month, err := cmd.Flags().GetInt("month")
		if err != nil {
			log.Fatal(err.Error())
		}

		day, err := cmd.Flags().GetInt("day")
		if err != nil {
			log.Fatal(err.Error())
		}

		setY := year > 0
		setM := month > 0
		setD := day > 0

		if !setY || (!setM && setD) {
			log.Fatal(`invalid flag combinations.
The correct flag combinations are:
  - year
  - year, month
  - year, month, day`)
		}

		if setY && setM && setD {
			d, err := jpcal.GetDay(year, month, day)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(outCSV(d))
			return
		}

		if setY && setM {
			ds, err := jpcal.AllDaysYM(year, month)
			if err != nil {
				log.Fatal(err.Error())
			}
			for _, d := range ds {
				fmt.Println(outCSV(d))
			}
			return
		}

		ds, err := jpcal.AllDays(year)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, d := range ds {
			fmt.Println(outCSV(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().IntP("year", "y", 0, "target year")
	getCmd.Flags().IntP("month", "m", 0, "target month")
	getCmd.Flags().IntP("day", "d", 0, "target day")
}

func outCSV(d jpcal.Day) string {
	return strings.Join([]string{d.Str(), string(d.Type()), d.Description()}, ",")
}
