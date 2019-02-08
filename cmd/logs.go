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

	"k8s.io/api/core/v1"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs POD",
	Short: "view logs of a TFJob",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			for _, pod := range args {
				if err := RunTFJobLogsCommand(pod); err != nil {
					log.Exit(1)
				}
			}
		}
	},
}

func RunTFJobLogsCommand(podName string) error {
	InitTFJobClient()
	logs, err := clientSet.CoreV1().Pods(namespace).GetLogs(podName, &v1.PodLogOptions{}).Do().Raw()
	if err != nil {
		log.WithFields(log.Fields{
			"name":      podName,
			"namespace": namespace,
			"error":     err,
		}).Warn("failed to get pod logs for TFJob")
	} else {
		log.WithFields(log.Fields{
			"name":      podName,
			"namespace": namespace,
		}).Info("successfully listed pods for TFJob")
		log.Info(string(logs))
	}
	return nil
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
