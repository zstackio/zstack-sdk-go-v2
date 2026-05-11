// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TaskProgressInventoryView TaskProgress
type TaskProgressInventoryView struct {
	BaseInfoView
	BaseTimeView
	TaskUuid string `json:"taskUuid,omitempty"`
	TaskName string `json:"taskName,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
	Opaque interface{} `json:"opaque,omitempty"`
	Time int64 `json:"time,omitempty"`
	SubTasks []*TaskProgressInventoryView `json:"subTasks,omitempty"`
	Arguments string `json:"arguments,omitempty"`
	ProgressDetail LongJobProgressDetailView `json:"progressDetail,omitempty"`
}

