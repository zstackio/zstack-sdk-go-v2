// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EcsImageInventoryView EcsImage
type EcsImageInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"localImageUuid,omitempty"`
	rest string `json:"ecsImageId,omitempty"`
	rest string `json:"name,omitempty"`
	rest int64 `json:"ecsImageSize,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"platform,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"ossMd5Sum,omitempty"`
	rest string `json:"format,omitempty"`
	rest string `json:"osName,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

