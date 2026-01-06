// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OssUploadPartsInventoryView OssUploadParts
type OssUploadPartsInventoryView struct {
	Id int64 `json:"id,omitempty"`
	UploadId string `json:"uploadId,omitempty"`
	PartNumber int `json:"partNumber,omitempty"`
	Total int `json:"total,omitempty"`
	ETag string `json:"eTag,omitempty"`
	PartSize int64 `json:"partSize,omitempty"`
	PartCRC int64 `json:"partCRC,omitempty"`
	OssBucketUuid string `json:"ossBucketUuid,omitempty"`
	FileKey string `json:"fileKey,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

