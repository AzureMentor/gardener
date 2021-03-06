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

package packetbotanist

import (
	"encoding/base64"
	"github.com/gardener/gardener/pkg/operation/common"
)

// GenerateCSIConfig generates the configuration for CSI charts
func (b *PacketBotanist) GenerateCSIConfig() (map[string]interface{}, error) {
	conf := map[string]interface{}{
		"credential": map[string]interface{}{
			"apiToken":  base64.StdEncoding.EncodeToString(b.Shoot.Secret.Data[PacketAPIKey]),
			"projectID": base64.StdEncoding.EncodeToString(b.Shoot.Secret.Data[ProjectID]),
		},
		"kubernetesVersion": b.ShootVersion(),
		"enabled":           true,
	}

	return b.InjectShootShootImages(conf,
		common.CSIPluginPacketImageName,
		common.CSINodeDriverRegistrarImageName,
	)
}

// GenerateEtcdBackupConfig returns the etcd backup configuration for the etcd Helm chart.
func (b *PacketBotanist) GenerateEtcdBackupConfig() (map[string][]byte, error) {
	return map[string][]byte{}, nil
}

// DeployCloudSpecificControlPlane does any last minute updates
func (b *PacketBotanist) DeployCloudSpecificControlPlane() error {
	return nil
}
