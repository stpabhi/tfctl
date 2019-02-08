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
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/duration"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get TFJob",
	Short: "display details about a TFJob",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			for _, tfJob := range args {
				if err := RunGetTFJobDetailCommand(tfJob); err != nil {
					log.Exit(1)
				}
			}
		}
	},
}

func RunGetTFJobDetailCommand(name string) error {
	tfJobClient := InitTFJobClient()
	getOpts := metav1.GetOptions{}
	_, err := tfJobClient.Get(name, getOpts)
	if err != nil {
		log.WithFields(log.Fields{
			"name":      name,
			"namespace": namespace,
			"error":     err,
		}).Warn("cannot find TFJob")
	} else {
		pods, err := clientSet.CoreV1().Pods(namespace).List(metav1.ListOptions{
			LabelSelector: fmt.Sprintf("group_name=kubeflow.org,tf_job_name=%s", name),
		})
		if err != nil {
			log.WithFields(log.Fields{
				"name":      name,
				"namespace": namespace,
				"error":     err,
			}).Warn("failed to list pods for TFJob")
		} else {
			log.WithFields(log.Fields{
				"name":      name,
				"namespace": namespace,
			}).Info("successfully listed pods for TFJob")
			for _, pod := range pods.Items {
				log.WithFields(log.Fields{
					"name":   pod.Name,
					"status": pod.Status.Phase,
					"age":    duration.ShortHumanDuration(time.Now().Sub(pod.CreationTimestamp.Time)),
				}).Info()
			}
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)
}
