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
	"io/ioutil"

	"github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1beta1"
	"github.com/kubeflow/tf-operator/pkg/client/clientset/versioned/scheme"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit TFJob",
	Short: "submit a TFJob",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			if err := RunSubmitTFJobCommand(args); err != nil {
				log.Exit(1)
			}
		}
	},
}

func RunSubmitTFJobCommand(filePaths []string) error {
	InitTFJobClient()
	body, err := ioutil.ReadFile(filePaths[0])
	if err != nil {
		log.Fatal(err)
	}
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, _ := decode(body, nil, nil)

	tfJob := obj.(*v1beta1.TFJob)
	log.Info(obj.GetObjectKind().GroupVersionKind().Kind)

	_, err = clientSet.CoreV1().Namespaces().Get(tfJob.Namespace, metav1.GetOptions{})

	if errors.IsNotFound(err) {
		// If namespace doesn't exist we create it
		_, nsErr := clientSet.CoreV1().Namespaces().Create(&v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: tfJob.Namespace}})
		if nsErr != nil {
			log.WithFields(log.Fields{
				"name":      tfJob.Name,
				"namespace": tfJob.Namespace,
				"error":     nsErr,
			}).Warn("failed to create namespace for TFJob")
		}
	} else if err != nil {
		log.WithFields(log.Fields{
			"name":      tfJob.Name,
			"namespace": tfJob.Namespace,
			"error":     err,
		}).Warn("failed to deploy TFJob")
	}

	job, err := tfJobClient.Create(tfJob)
	if err != nil {
		log.WithFields(log.Fields{
			"name":      tfJob.Name,
			"namespace": tfJob.Namespace,
			"error":     err,
		}).Warn("failed to deploy TFJob")
	} else {
		log.WithFields(log.Fields{
			"name":      job.Name,
			"namespace": job.Namespace,
		}).Warn("successfully deployed TFJob")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(submitCmd)
}
