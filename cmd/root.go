/*
Copyright © 2022 Utkarsh Srivastava <utkarsh@sagacious.dev>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/utkarsh-pro/ucs/pkg/chtsh"
	"github.com/utkarsh-pro/ucs/pkg/qparser"
	"github.com/utkarsh-pro/ucs/pkg/utils"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ucs",
	Short: "ucs is a very thin CLI wrapper for cht.sh",
	Run: func(cmd *cobra.Command, args []string) {
		result := qparser.Parse(args)
		data, err := chtsh.New(result).Find()
		utils.LogIfError(err)

		if err == nil {
			fmt.Println(data)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
