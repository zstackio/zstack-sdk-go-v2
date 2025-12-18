// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MonitorGroupInstanceInventoryView MonitorGroupInstance
type MonitorGroupInstanceInventoryView struct {
	GroupUuid string `json:"groupUuid,omitempty"`
	InstanceResourceType string `json:"instanceResourceType,omitempty"`
	InstanceUuid string `json:"instanceUuid,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

