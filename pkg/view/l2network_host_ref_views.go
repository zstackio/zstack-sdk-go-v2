// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L2NetworkHostRefInventoryView L2NetworkHostRef
type L2NetworkHostRefInventoryView struct {
	HostUuid       string    `json:"hostUuid,omitempty"`
	L2NetworkUuid  string    `json:"l2NetworkUuid,omitempty"`
	L2ProviderType string    `json:"l2ProviderType,omitempty"`
	AttachStatus   string    `json:"attachStatus,omitempty"`
	CreateDate     time.Time `json:"createDate,omitempty"`
	LastOpDate     time.Time `json:"lastOpDate,omitempty"`
}
