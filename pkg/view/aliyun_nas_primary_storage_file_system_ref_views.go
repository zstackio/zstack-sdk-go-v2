// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunNasPrimaryStorageFileSystemRefInventoryView AliyunNasPrimaryStorageFileSystemRef
type AliyunNasPrimaryStorageFileSystemRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"nasFileSystemUuid,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

