// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ImageCacheInventoryView ImageCache
type ImageCacheInventoryView struct {
	Id int64 `json:"id,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	ImageUuid *string `json:"imageUuid,omitempty"`
	InstallUrl *string `json:"installUrl,omitempty"`
	MediaType *string `json:"mediaType,omitempty"`
	Size int64 `json:"size,omitempty"`
	Md5sum *string `json:"md5sum,omitempty"`
	State *string `json:"state,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryImageCacheView QueryImageCache
type QueryImageCacheView struct {
	Inventories []ImageCacheInventoryView `json:"inventories,omitempty"`
}

