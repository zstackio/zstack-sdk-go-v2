// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaremetalImageCacheInventoryView BaremetalImageCache
type BaremetalImageCacheInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"pxeServerUuid,omitempty"`
	rest string `json:"imageUuid,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"installUrl,omitempty"`
	rest string `json:"mediaType,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest int64 `json:"actualSize,omitempty"`
	rest string `json:"md5sum,omitempty"`
	rest int64 `json:"utilization,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

