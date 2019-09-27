/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package list

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alandtsang/autocompletion/internal/common"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list directory contents",
	Long: `List information about the FILEs (the current directory by default).

Sort entries alphabetically if none of -cftuvSUX nor --sort is specified.`,
	Run: func(cmd *cobra.Command, args []string) {
		getArgs := cmd.Flags().Args()
		do(getArgs)
	},
}

// Init Initialize list command.
func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "list all")
}

func do(names []string) {
	for _, name := range names {
		list(name)
		fmt.Println()
	}
}

func list(name string) {
	isDir, err := common.JudgeType(name)
	if err != nil {
		return
	}
	if isDir {
		listDir(name)
	} else {
		listFile(name)
	}
}

func listFile(name string) {
	fmt.Println(name)
}

func listDir(dir string) {
	fmt.Printf("%s:\n", dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Printf("%s  ", f.Name())
	}
	fmt.Println()
}
