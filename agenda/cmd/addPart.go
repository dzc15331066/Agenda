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

	"github.com/spf13/cobra"
)

// addPartCmd represents the addPart command
var addPartCmd = &cobra.Command{
	Use:   "addPart",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addPart called")
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringSlice("part")
		if as.AgendaStorage.CurUser == (entity.User{}) {
			fmt.Println("[error]: not registered yet!")
		} else {
			for _, username := range participators {
				res := as.AddParticipator(username, title)
				if !res {
					fmt.Printf("[error]: %s hasn't registered!\n", username)
				} else {
					fmt.Println("[success]: all the participators are added.")
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(addPartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addPartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addPartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addPartCmd.Flags().StringP("title", "t", "", "user --title or -t [meeting's title]")
	addPartCmd.Flags().StringSliceP("part", "p", make([]string, 0), "user --part or -p [participators]")
}
