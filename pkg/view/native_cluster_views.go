// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NativeClusterInventoryView NativeCluster
type NativeClusterInventoryView struct {
	BizUrl string `json:"bizUrl,omitempty"`
	MasterUrl string `json:"masterUrl,omitempty"`
	KubeConfig string `json:"kubeConfig,omitempty"`
	PrometheusURL string `json:"prometheusURL,omitempty"`
	Version string `json:"version,omitempty"`
	NodeCount int `json:"nodeCount,omitempty"`
	CreateType string `json:"createType,omitempty"`
	Status string `json:"status,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

