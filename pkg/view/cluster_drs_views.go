// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ClusterDRSInventoryView ClusterDRS
type ClusterDRSInventoryView struct {
	ClusterUuid string `json:"clusterUuid,omitempty"`
	State string `json:"state,omitempty"`
	BalancedState string `json:"balancedState,omitempty"`
	AutomationLevel string `json:"automationLevel,omitempty"`
	Thresholds []interface{} `json:"thresholds,omitempty"`
	ThresholdDuration int `json:"thresholdDuration,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

