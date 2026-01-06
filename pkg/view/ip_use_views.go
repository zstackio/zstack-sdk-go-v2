// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IpUseInventoryView IpUse
type IpUseInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	UsedIpUuid string `json:"usedIpUuid,omitempty"`
	ServiceId string `json:"serviceId,omitempty"`
	Use string `json:"use,omitempty"`
	Details string `json:"details,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

