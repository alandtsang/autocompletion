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
package cat

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alandtsang/autocompletion/internal/common"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "concatenate files and print on the standard output",
	Long:  `Concatenate FILE(s) to standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		getArgs := cmd.Flags().Args()
		validate(getArgs)
		do(getArgs)
	},
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(catCmd)
}

func validate(args []string) {
	argsLen := len(args)
	if argsLen == 0 {
		log.Fatalln("please input file")
	}
}

func do(names []string) {
	for _, name := range names {
		cat(name)
		fmt.Println()
	}
}

// Init Initialize cat command.
func cat(name string) {
	isDir, err := common.JudgeType(name)
	if err != nil {
		log.Fatalf("cat %s failed, %s\n", name, err.Error())
	}

	if isDir {
		fmt.Printf("cat: %s Is a directory\n", name)
		return
	}

	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("cat %s failed, %s\n", name, err.Error())
	}
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatalf("cat %s failed, %s\n", name, err.Error())
	}
}
