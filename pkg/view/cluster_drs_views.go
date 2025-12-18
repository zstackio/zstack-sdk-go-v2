// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ClusterDRSInventoryView ClusterDRS
type ClusterDRSInventoryView struct {
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"balancedState,omitempty"`
	rest string `json:"automationLevel,omitempty"`
	rest []interface{} `json:"thresholds,omitempty"`
	rest int `json:"thresholdDuration,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
}

