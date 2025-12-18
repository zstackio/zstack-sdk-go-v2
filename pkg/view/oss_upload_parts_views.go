// Copyright (c) ZStack.io, Inc.

package view

import "time"

// OssUploadPartsInventoryView OssUploadParts
type OssUploadPartsInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"uploadId,omitempty"`
	rest int `json:"partNumber,omitempty"`
	rest int `json:"total,omitempty"`
	rest string `json:"eTag,omitempty"`
	rest int64 `json:"partSize,omitempty"`
	rest int64 `json:"partCRC,omitempty"`
	rest string `json:"ossBucketUuid,omitempty"`
	rest string `json:"fileKey,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

