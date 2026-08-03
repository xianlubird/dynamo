/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package dynamo

import (
	"github.com/ai-dynamo/dynamo/deploy/operator/api/v1beta1"
	runtimefeatures "github.com/ai-dynamo/dynamo/deploy/operator/internal/features/runtime"
	"github.com/ai-dynamo/dynamo/deploy/operator/internal/runtimeversion"
)

// runtimeProfileForComponent returns the runtime profile for a component.
func runtimeProfileForComponent(component *v1beta1.DynamoComponentDeploymentSharedSpec) runtimefeatures.RuntimeProfile {
	if component == nil {
		return runtimefeatures.RuntimeProfile{}
	}

	image := ""
	if main := GetMainContainer(component); main != nil {
		image = main.Image
	}

	version, err := runtimeversion.Resolve(image, component.RuntimeVersionOverride)
	if err != nil {
		return runtimefeatures.RuntimeProfile{}
	}

	return runtimefeatures.ProfileForVersion(&version)
}
