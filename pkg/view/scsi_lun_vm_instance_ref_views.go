// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ScsiLunVmInstanceRefInventoryView ScsiLunVmInstanceRef
type ScsiLunVmInstanceRefInventoryView struct {
	rest string `json:"scsiLunUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest int `json:"deviceId,omitempty"`
	rest bool `json:"attachMultipath,omitempty"`
}

