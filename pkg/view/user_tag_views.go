// Copyright (c) ZStack.io, Inc.

package view

import "time"

// UserTagInventoryView UserTag
type UserTagInventoryView struct {
	rest string `json:"tagPatternUuid,omitempty"`
	rest TagPatternInventoryView `json:"tagPattern,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"tag,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

