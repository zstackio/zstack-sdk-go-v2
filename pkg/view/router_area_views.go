// Copyright (c) ZStack.io, Inc.

package view

import "time"

// RouterAreaInventoryView RouterArea
type RouterAreaInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"areaId,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"authentication,omitempty"`
	rest string `json:"password,omitempty"`
	rest int `json:"keyId,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

