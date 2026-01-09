// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ZBoxLocationRefInventoryView ZBoxLocationRef
type ZBoxLocationRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ZboxUuid *string `json:"zboxUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

