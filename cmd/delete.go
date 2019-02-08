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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete TFJob",
	Short: "delete a TFJob and its associated pods",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			for _, tfJob := range args {
				if err := RunDeleteTFJobCommand(tfJob); err != nil {
					log.Exit(1)
				}
			}
		}
	},
}

func RunDeleteTFJobCommand(name string) error {
	tfJobClient := InitTFJobClient()
	deleteOpts := metav1.DeleteOptions{}
	err := tfJobClient.Delete(name, &deleteOpts)
	if err != nil {
		log.WithFields(log.Fields{
			"name":      name,
			"namespace": namespace,
			"error":     err,
		}).Warn("failed to delete TFJob")
	} else {
		log.WithFields(log.Fields{
			"name":      name,
			"namespace": namespace,
		}).Info("successfully deleted TFJob")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
