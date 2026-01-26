// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OssBucketPropertyView OssBucketProperty
type OssBucketPropertyView struct {
	BucketName string `json:"bucketName,omitempty"`
	RegionId string `json:"regionId,omitempty"`
}

