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

package azurebotanist

import (
	"github.com/gardener/gardener/pkg/operation/common"
)

// GenerateCSIConfig generates the configuration for CSI charts
func (b *AzureBotanist) GenerateCSIConfig() (map[string]interface{}, error) {
	return nil, nil
}

// GenerateEtcdBackupConfig returns the etcd backup configuration for the etcd Helm chart.
func (b *AzureBotanist) GenerateEtcdBackupConfig() (map[string][]byte, error) {
	var (
		storageAccountName = "storageAccountName"
		storageAccessKey   = "storageAccessKey"
		containerName      = "containerName"
	)
	tf, err := b.NewBackupInfrastructureTerraformer()
	if err != nil {
		return nil, err
	}
	stateVariables, err := tf.GetStateOutputVariables(storageAccountName, storageAccessKey, containerName)
	if err != nil {
		return nil, err
	}

	secretData := map[string][]byte{
		common.BackupBucketName: []byte(stateVariables[containerName]),
		"storage-account":       []byte(stateVariables[storageAccountName]),
		"storage-key":           []byte(stateVariables[storageAccessKey]),
	}

	return secretData, nil
}

// DeployCloudSpecificControlPlane does currently nothing for Azure.
func (b *AzureBotanist) DeployCloudSpecificControlPlane() error {
	return nil
}
