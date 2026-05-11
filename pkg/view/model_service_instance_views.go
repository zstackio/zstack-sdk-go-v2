// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ModelServiceInstanceInventoryView ModelServiceInstance
type ModelServiceInstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
	ModelServiceGroupUuid string `json:"modelServiceGroupUuid,omitempty"`
	Yaml string `json:"yaml,omitempty"`
	K8sResourceYaml string `json:"k8sResourceYaml,omitempty"`
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
	UrlMaps map[string]string `json:"urlMaps,omitempty"`
	InternalUrl string `json:"internalUrl,omitempty"`
	JupyterUrl string `json:"jupyterUrl,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	NodeRank int `json:"nodeRank,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	GpuVendor string `json:"gpuVendor,omitempty"`
	Vm VmInstanceInventoryView `json:"vm,omitempty"`
}

