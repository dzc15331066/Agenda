// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"Agenda/agenda/entity"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cm called")
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetStringSlice("part")
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		sponsor := as.AgendaStorage.CurUser.Name
		// time parse
		if as.AgendaStorage.CurUser == (entity.User{}) {
			fmt.Println("[error]: not login yet!")
		} else {
			const shortForm = "2006-Jan-02"
			s, _ := time.Parse(shortForm, start)
			e, _ := time.Parse(shortForm, end)
			res := as.AddMeeting(sponsor, title, s, e, participator)
			if !res {
				fmt.Println("[error]: Adding meeting fails!check out and try again")
			} else {
				fmt.Println("[success]: Adding meeting successfully!")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(cmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmCmd.Flags().StringP("title", "t", "", "use -title [meeting's title] or -t [meeting's title]")
	cmCmd.Flags().StringSliceP("part", "p", make([]string, 1), "use -part [participators] or -p [participators]")
	cmCmd.Flags().StringP("start", "s", "2017-10-25", "use -start or -s [year-month-day]")
	cmCmd.Flags().StringP("end", "e", "2017-10-25", "use -end or -e [year-month-day]")
}
