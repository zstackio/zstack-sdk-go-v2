// Copyright (c) ZStack.io, Inc.

package view

import "time"

// OssBucketDomainInventoryView OssBucketDomain
type OssBucketDomainInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"ossBucketUuid,omitempty"`
	rest string `json:"ossDomain,omitempty"`
	rest string `json:"ossKey,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

