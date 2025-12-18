// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PortMirrorSessionInventoryView PortMirrorSession
type PortMirrorSessionInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	InternalId int64 `json:"internalId,omitempty"`
	SrcEndPoint string `json:"srcEndPoint,omitempty"`
	Type string `json:"type,omitempty"`
	DstEndPoint string `json:"dstEndPoint,omitempty"`
	PortMirrorUuid string `json:"portMirrorUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

