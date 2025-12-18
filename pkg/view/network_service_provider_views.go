// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NetworkServiceProviderInventoryView NetworkServiceProvider
type NetworkServiceProviderInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	NetworkServiceTypes []string `json:"networkServiceTypes,omitempty"`
	AttachedL2NetworkUuids []string `json:"attachedL2NetworkUuids,omitempty"`
}

