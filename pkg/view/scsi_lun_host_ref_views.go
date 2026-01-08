// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ScsiLunHostRefInventoryView ScsiLunHostRef
type ScsiLunHostRefInventoryView struct {
	ScsiLunUuid string    `json:"scsiLunUuid,omitempty"`
	HostUuid    string    `json:"hostUuid,omitempty"`
	Path        string    `json:"path,omitempty"`
	Hctl        string    `json:"hctl,omitempty"`
	CreateDate  time.Time `json:"createDate,omitempty"`
	LastOpDate  time.Time `json:"lastOpDate,omitempty"`
}
