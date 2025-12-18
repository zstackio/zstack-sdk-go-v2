// Copyright (c) ZStack.io, Inc.

package view

import "time"

// UserProxyConfigResourceRefInventoryView UserProxyConfigResourceRef
type UserProxyConfigResourceRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"proxyUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

