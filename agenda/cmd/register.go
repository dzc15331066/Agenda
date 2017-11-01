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
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("contact")
		log.Infoln("Register:")
		err := as.UserRegister(username, password, email, phone)
		message(err, "[success]:OK! register successfully!")
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("user", "u", "", "Help message for username")
	registerCmd.Flags().StringP("password", "p", "", "Help message for password")
	registerCmd.Flags().StringP("email", "e", "", "Help message for email")
	registerCmd.Flags().StringP("contact", "c", "", "Help message for phone")
}
