// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VniRangeInventoryView VniRange
type VniRangeInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StartVni int `json:"startVni,omitempty"`
	EndVni int `json:"endVni,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid,omitempty"`
}

