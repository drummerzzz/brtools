/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/drummerzzz/cpftools"
	"github.com/drummerzzz/cpftools/src/utils"
	"github.com/spf13/cobra"
)

// cpfCmd represents the cpf command
var cpfCmd = &cobra.Command{
	Use:   "cpf",
	Short: "CPF utils",
	Long:  `Generate and validate CPF:`,

	Run: func(cmd *cobra.Command, args []string) {
		var cpf string
		cpf, _ = cmd.Flags().GetString("validate")
		if cpf != "" {
			fmt.Println(utils.GetOnlyDigits(cpf), cpftools.IsValid(cpf))
			return
		}
		isClean, _ := cmd.Flags().GetBool("clean")
		if isClean {
			cpf = cpftools.GenerateWithoutMask()
		} else {
			cpf = cpftools.GenerateWithMask()
		}
		fmt.Println(cpf)
	},
}

func init() {
	cpfCmd.Flags().BoolP("clean", "c", false, "Generate a cpf without mask")
	cpfCmd.Flags().StringP("validate", "v", "", `Validate a cpf`)
	rootCmd.AddCommand(cpfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
