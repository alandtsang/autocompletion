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
package completion

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var validArgs = []string{"pod", "node", "service", "replicationcontroller"}
var shells = []string{}

var (
	completionShells = map[string]func(out io.Writer, cmd *cobra.Command) error{
		"bash": runCompletionBash,
		//"zsh":  runCompletionZsh,
	}
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Output shell completion code for the specified shell (bash or zsh)",
	Long: `Output shell completion code for the specified shell (bash or zsh). The shell code must be evaluated to provide
interactive completion of kubectl commands.

This can be done by sourcing it from the .bash_profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := RunCompletion(os.Stdout, cmd, args)
		if err != nil {
			log.Fatal(err)
			return
		}
	},
	ValidArgs: shells,
}

func Init(rootCmd *cobra.Command) {
	for s := range completionShells {
		shells = append(shells, s)
	}
	rootCmd.AddCommand(completionCmd)
}

// RunCompletion checks given arguments and executes command.
func RunCompletion(out io.Writer, cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Shell not specified.")
	}
	if len(args) > 1 {
		return errors.New("Too many arguments. Expected only the shell type.")
	}
	run, found := completionShells[args[0]]
	if !found {
		return errors.New("Unsupported shell type")
	}

	return run(out, cmd.Parent())
}

func runCompletionBash(out io.Writer, kubectl *cobra.Command) error {
	return kubectl.GenBashCompletion(out)
}
