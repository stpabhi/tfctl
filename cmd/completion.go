// Copyright © 2019 Abhilash Pallerlamudi <stp.abhi@gmail.com>
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
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	bashCompletionFunc = `
__tfctl_get_tfjobs() {
	local tfctl_out
	if tfctl_out=$(tfctl list --output name 2>/dev/null); then
		COMPREPLY+=( $( compgen -W "${tfctl_out[*]}" -- "$cur" ) )
	fi
}

__tfctl_custom_func() {
	case ${last_command} in
		tfctl_delete | tfctl_get | tfctl_logs |\
		tfctl_resubmit | tfctl_resume | tfctl_retry | tfctl_suspend |\
		tfctl_terminate | tfctl_wait | tfctl_watch)
			__tfctl_get_tfjobs
			return
			;;
		*)
			;;
	esac
}
	`
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion SHELL",
	Short: "output shell completion code for the specified shell (bash or zsh)",
	Long: `Write bash or zsh shell completion code to standard output.

For bash, ensure you have bash completions installed and enabled.
To access completions in your current shell, run
$ source <(tfctl completion bash)
Alternatively, write it to a file and source in .bash_profile

For zsh, output to a file in a directory referenced by the $fpath shell
variable.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.HelpFunc()(cmd, args)
			os.Exit(1)
		}
		shell := args[0]
		rootCmd.BashCompletionFunction = bashCompletionFunc
		availableCompletions := map[string]func(io.Writer) error{
			"bash": rootCmd.GenBashCompletion,
			"zsh":  rootCmd.GenZshCompletion,
		}
		completion, ok := availableCompletions[shell]
		if !ok {
			fmt.Printf("Invalid shell '%s'. The supported shells are bash and zsh.\n", shell)
			os.Exit(1)
		}
		if err := completion(os.Stdout); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
