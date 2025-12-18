// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NetworkRouterAreaRefInventoryView NetworkRouterAreaRef
type NetworkRouterAreaRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	ApplianceVmType string `json:"applianceVmType,omitempty"`
	RouterAreaUuid string `json:"routerAreaUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

