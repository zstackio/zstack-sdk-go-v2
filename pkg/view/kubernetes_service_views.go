// Copyright (c) ZStack.io, Inc.

package view

import "time"

// KubernetesServiceInventoryView KubernetesService
type KubernetesServiceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"clusterIp,omitempty"`
	rest string `json:"externalIp,omitempty"`
	rest string `json:"ports,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
	rest int64 `json:"clusterId,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

