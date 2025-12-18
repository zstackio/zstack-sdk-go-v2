// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OssBucketInventoryView OssBucket
type OssBucketInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	BucketName string `json:"bucketName,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	Current string `json:"current,omitempty"`
	RegionName string `json:"regionName,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

