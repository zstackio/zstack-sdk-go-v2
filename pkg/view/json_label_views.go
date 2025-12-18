// Copyright (c) ZStack.io, Inc.

package view

import "time"

// JsonLabelInventoryView JsonLabel
type JsonLabelInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"labelKey,omitempty"`
	rest string `json:"labelValue,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

