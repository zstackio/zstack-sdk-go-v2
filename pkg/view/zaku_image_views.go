// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ZakuImageInventoryView ZakuImage
type ZakuImageInventoryView struct {
	Name string `json:"name,omitempty"`
	TagCount int `json:"tagCount,omitempty"`
	PullCount int `json:"pullCount,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

