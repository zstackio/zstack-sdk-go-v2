// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasPrimaryStorageFileSystemRefInventoryView AliyunNasPrimaryStorageFileSystemRef
type AliyunNasPrimaryStorageFileSystemRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	NasFileSystemUuid string `json:"nasFileSystemUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

