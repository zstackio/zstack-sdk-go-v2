// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NfvInstGroupL3NetworkRefInventoryView NfvInstGroupL3NetworkRef
type NfvInstGroupL3NetworkRefInventoryView struct {
	NfvInstGroupUuid string `json:"nfvInstGroupUuid,omitempty"`
	NetworkServiceUuid string `json:"networkServiceUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	L3NetworkCategory string `json:"l3NetworkCategory,omitempty"`
	L3NetworkType string `json:"l3NetworkType,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

