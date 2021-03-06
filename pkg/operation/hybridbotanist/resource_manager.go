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

package hybridbotanist

import (
	"context"
	"path/filepath"

	"github.com/gardener/gardener/pkg/operation/common"
)

// DeployGardenerResourceManager deploys the gardener-resource-manager which will use CRD resources in order
// to ensure that they exist in a cluster/reconcile them in case somebody changed something.
func (b *HybridBotanist) DeployGardenerResourceManager(ctx context.Context) error {
	var name = "gardener-resource-manager"

	defaultValues := map[string]interface{}{
		"podAnnotations": map[string]interface{}{
			"checksum/secret-" + name: b.CheckSums[name],
		},
		"replicas": b.Shoot.GetReplicas(1),
	}

	values, err := b.InjectSeedShootImages(defaultValues, common.GardenerResourceManagerImageName)
	if err != nil {
		return err
	}

	return b.ApplyChartSeed(filepath.Join(common.ChartPath, "seed-controlplane", "charts", name), b.Shoot.SeedNamespace, name, values, nil)
}
