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
	"github.com/kubeflow/tf-operator/pkg/client/clientset/versioned"
	"github.com/kubeflow/tf-operator/pkg/client/clientset/versioned/typed/kubeflow/v1beta1"
	log "github.com/sirupsen/logrus"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	namespace    string
	restConfig   *rest.Config
	clientConfig clientcmd.ClientConfig
	clientSet    *kubernetes.Clientset
	tfJobClient  v1beta1.TFJobInterface
)

func initKubeClient() *kubernetes.Clientset {
	if clientSet != nil {
		return clientSet
	}
	var err error
	restConfig, err = clientConfig.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientSet, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatal(err)
	}

	return clientSet
}

func InitTFJobClient(ns ...string) v1beta1.TFJobInterface {
	if tfJobClient != nil {
		return tfJobClient
	}
	initKubeClient()
	var err error
	if len(ns) > 0 {
		namespace = ns[0]
	} else {
		namespace, _, err = clientConfig.Namespace()
		if err != nil {
			log.Fatal(err)
		}
	}
	tfClientset := versioned.NewForConfigOrDie(restConfig)
	tfJobClient = tfClientset.KubeflowV1beta1().TFJobs(namespace)
	return tfJobClient
}
