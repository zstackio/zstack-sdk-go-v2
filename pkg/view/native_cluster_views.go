// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NativeClusterInventoryView NativeCluster
type NativeClusterInventoryView struct {
	rest string `json:"bizUrl,omitempty"`
	rest string `json:"masterUrl,omitempty"`
	rest string `json:"kubeConfig,omitempty"`
	rest string `json:"prometheusURL,omitempty"`
	rest string `json:"version,omitempty"`
	rest int `json:"nodeCount,omitempty"`
	rest string `json:"createType,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"hypervisorType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"architecture,omitempty"`
}

