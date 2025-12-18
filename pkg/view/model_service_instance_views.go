// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ModelServiceInstanceInventoryView ModelServiceInstance
type ModelServiceInstanceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"modelServiceGroupUuid,omitempty"`
	rest string `json:"yaml,omitempty"`
	rest string `json:"k8sResourceYaml,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"url,omitempty"`
	rest map[string]string `json:"urlMaps,omitempty"`
	rest string `json:"internalUrl,omitempty"`
	rest string `json:"jupyterUrl,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest int `json:"nodeRank,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest VmInstanceInventoryView `json:"vm,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

