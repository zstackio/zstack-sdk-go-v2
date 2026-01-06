// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasPrimaryStorageFileSystemRefInventoryView AliyunNasPrimaryStorageFileSystemRef
type AliyunNasPrimaryStorageFileSystemRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	NasFileSystemUuid string `json:"nasFileSystemUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

