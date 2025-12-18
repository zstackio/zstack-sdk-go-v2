// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IscsiLunInventoryView IscsiLun
type IscsiLunInventoryView struct {
	rest string `json:"iscsiTargetUuid,omitempty"`
	rest []ScsiLunHostRefInventoryView `json:"scsiLunHostRefs,omitempty"`
	rest []ScsiLunVmInstanceRefInventoryView `json:"scsiLunVmInstanceRefs,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"wwid,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest string `json:"model,omitempty"`
	rest string `json:"wwn,omitempty"`
	rest string `json:"serial,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"hctl,omitempty"`
	rest string `json:"path,omitempty"`
	rest string `json:"state,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"multipathDeviceUuid,omitempty"`
	rest string `json:"source,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

