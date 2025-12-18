// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SlbGroupInventoryView SlbGroup
type SlbGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	BackendType string `json:"backendType,omitempty"`
	DeployType string `json:"deployType,omitempty"`
	SlbOfferingUuid string `json:"slbOfferingUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ConfigVersion int64 `json:"configVersion,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	SlbVms []SlbVmInstanceInventoryView `json:"slbVms,omitempty"`
	Lbs []SlbLoadBalancerInventoryView `json:"lbs,omitempty"`
	Networks []SlbGroupL3NetworkRefInventoryView `json:"networks,omitempty"`
	MonitorIps []SlbGroupMonitorIpInventoryView `json:"monitorIps,omitempty"`
}

