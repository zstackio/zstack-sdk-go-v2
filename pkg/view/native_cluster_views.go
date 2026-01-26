// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NativeClusterInventoryView NativeCluster
type NativeClusterInventoryView struct {
	BaseInfoView
	BaseTimeView
	BizUrl string `json:"bizUrl,omitempty"`
	MasterUrl string `json:"masterUrl,omitempty"`
	KubeConfig string `json:"kubeConfig,omitempty"`
	PrometheusURL string `json:"prometheusURL,omitempty"`
	Version string `json:"version,omitempty"`
	NodeCount int `json:"nodeCount,omitempty"`
	CreateType string `json:"createType,omitempty"`
	Status string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

// QueryNativeClusterView QueryNativeCluster
type QueryNativeClusterView struct {
	Inventories []NativeClusterInventoryView `json:"inventories,omitempty"`
}

