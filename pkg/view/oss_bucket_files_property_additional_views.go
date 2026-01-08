// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OssBucketFilesPropertyView OssBucketFilesProperty
type OssBucketFilesPropertyView struct {
	RegionId  string   `json:"regionId,omitempty"`
	OssBucket string   `json:"ossBucket,omitempty"`
	Files     []string `json:"files,omitempty"`
}
