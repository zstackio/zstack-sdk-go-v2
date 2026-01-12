// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OssBucketDomainInventoryView OssBucketDomain
type OssBucketDomainInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	OssBucketUuid *string `json:"ossBucketUuid,omitempty"`
	OssDomain *string `json:"ossDomain,omitempty"`
	OssKey *string `json:"ossKey,omitempty"`
}

