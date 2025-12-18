// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ModelCenterInventoryView ModelCenter
type ModelCenterInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"parameters,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest int `json:"managementPort,omitempty"`
	rest string `json:"storageNetworkUuid,omitempty"`
	rest string `json:"serviceNetworkUuid,omitempty"`
	rest string `json:"containerRegistry,omitempty"`
	rest string `json:"containerStorageNetwork,omitempty"`
	rest string `json:"containerNetwork,omitempty"`
	rest ModelCenterCapacityInventoryView `json:"capacity,omitempty"`
	rest ZdfsInventoryView `json:"zdfs,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

