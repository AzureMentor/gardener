// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alicloudbotanist

import (
	"github.com/gardener/gardener/pkg/operation/common"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeployKube2IAMResources - Not needed on Alicloud
func (b *AlicloudBotanist) DeployKube2IAMResources() error {
	return nil
}

// DestroyKube2IAMResources - Not needed on Alicloud
func (b *AlicloudBotanist) DestroyKube2IAMResources() error {
	return nil
}

// GenerateKube2IAMConfig - Not needed on Alicloud.
func (b *AlicloudBotanist) GenerateKube2IAMConfig() (map[string]interface{}, error) {
	return common.GenerateAddonConfig(nil, false), nil
}

// GenerateNginxIngressConfig generates values which are required to render the chart nginx-ingress properly.
func (b *AlicloudBotanist) GenerateNginxIngressConfig() (map[string]interface{}, error) {
	return common.GenerateAddonConfig(map[string]interface{}{
		"controller": map[string]interface{}{
			"config": map[string]interface{}{
				"use-proxy-protocol": "false",
			},
		},
	}, b.Shoot.NginxIngressEnabled()), nil
}

// GenerateStorageClassesConfig generates values which are required to render the chart shoot-storageclasses properly.
func (b *AlicloudBotanist) GenerateStorageClassesConfig() (map[string]interface{}, error) {
	//TODO Delete in next release. Add encrypted flag in storge class
	storageclassList, err := b.Operation.K8sShootClient.Kubernetes().StorageV1().StorageClasses().List(metav1.ListOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		return nil, err
	}
	for _, sc := range storageclassList.Items {
		if sc.Provisioner == "diskplugin.csi.alibabacloud.com" && (sc.Name == "default" || sc.Name == "gardener.cloud-fast") {
			if _, ok := sc.Parameters["encrypted"]; !ok {
				if deleteErr := b.Operation.K8sShootClient.Kubernetes().StorageV1().StorageClasses().Delete(sc.Name, &metav1.DeleteOptions{}); deleteErr != nil {
					return nil, err
				}
			}
		}
	}

	return map[string]interface{}{
		"StorageClasses": []map[string]interface{}{
			{
				"Name":           "default",
				"IsDefaultClass": true,
				"Provisioner":    "diskplugin.csi.alibabacloud.com",
				"Parameters": map[string]interface{}{
					"csi.storage.k8s.io/fstype": "ext4",
					"type":                      "cloud_ssd",
					"readOnly":                  "false",
					"encrypted":                 "true",
				},
			},
		},
	}, nil

}

// GenerateVPNShootConfig generate cloud-specific vpn override - nothing unique for alicloud
func (b *AlicloudBotanist) GenerateVPNShootConfig() (map[string]interface{}, error) {
	return nil, nil
}
