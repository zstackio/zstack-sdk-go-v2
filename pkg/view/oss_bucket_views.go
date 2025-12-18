// Copyright (c) ZStack.io, Inc.

package view

import "time"

// OssBucketInventoryView OssBucket
type OssBucketInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"bucketName,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"current,omitempty"`
	rest string `json:"regionName,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

