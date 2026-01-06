// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ResourceStackVmPortRefInventoryView ResourceStackVmPortRef
type ResourceStackVmPortRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	StackUuid string `json:"stackUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	Port int `json:"port,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

