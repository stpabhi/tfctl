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
	"github.com/kubeflow/tf-operator/pkg/client/clientset/versioned/typed/kubeflow/v1beta1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var listArgs listFlags

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list TFJobs",
	Run: func(cmd *cobra.Command, args []string) {
		if err := RunListTFJobCommand(); err != nil {
			log.Exit(1)
		}
	},
}

type listFlags struct {
	allNamespaces bool
	status        string
	completed     bool
	running       bool
	output        string
	since         string
}

func RunListTFJobCommand() error {
	var tfClient v1beta1.TFJobInterface
	if listArgs.allNamespaces {
		tfClient = InitTFJobClient(apiv1.NamespaceAll)
	} else {
		tfClient = InitTFJobClient()
	}
	listOpts := metav1.ListOptions{}
	tfJobList, err := tfClient.List(listOpts)
	if err != nil {
		log.WithFields(log.Fields{
			"namespace": namespace,
			"error":     err,
		}).Warn("failed to list TFJobs")
	} else {
		if len(tfJobList.Items) == 0 {
			log.WithFields(log.Fields{
				"namespace": namespace,
			}).Info("no TFJobs to list")
		} else {
			for _, tfjob := range tfJobList.Items {
				log.Info(tfjob.Name)
			}
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&listArgs.allNamespaces, "all-namespaces", false, "Show tensorflow jobs from all namespaces")
}
