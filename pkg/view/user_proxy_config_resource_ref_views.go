// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserProxyConfigResourceRefInventoryView UserProxyConfigResourceRef
type UserProxyConfigResourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id *int64 `json:"id,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	ProxyUuid *string `json:"proxyUuid,omitempty"`
}

// AddProxyToResourceEventView AddProxyToResourceEvent
type AddProxyToResourceEventView struct {
	Inventory UserProxyConfigResourceRefInventoryView `json:"inventory,omitempty"`
}

