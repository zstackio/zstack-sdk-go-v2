// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ConnectionAccessPointInventoryView ConnectionAccessPoint
type ConnectionAccessPointInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"accessPointId,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"hostOperator,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

