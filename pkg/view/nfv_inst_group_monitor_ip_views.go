// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NfvInstGroupMonitorIpInventoryView NfvInstGroupMonitorIp
type NfvInstGroupMonitorIpInventoryView struct {
	Id *int64 `json:"id,omitempty"`
	NfvInstGroupUuid *string `json:"nfvInstGroupUuid,omitempty"`
	MonitorIp *string `json:"monitorIp,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

