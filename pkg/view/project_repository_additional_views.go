// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ProjectRepositoryInventoryView ProjectRepository
type ProjectRepositoryInventoryView struct {
	ZeProjectID int    `json:"zeProjectID,omitempty"`
	Type        string `json:"type,omitempty"`
	Readonly    bool   `json:"readonly,omitempty"`
	Name        string `json:"name,omitempty"`
	ImageCount  int    `json:"imageCount,omitempty"`
	ID          int64  `json:"ID,omitempty"`
	Desc        string `json:"desc,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	ChartCount  int    `json:"chartCount,omitempty"`
}
