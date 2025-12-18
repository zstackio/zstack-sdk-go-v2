// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ImageCacheInventoryView ImageCache
type ImageCacheInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"imageUuid,omitempty"`
	rest string `json:"installUrl,omitempty"`
	rest string `json:"mediaType,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"md5sum,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

