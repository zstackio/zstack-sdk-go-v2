// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NfvInstGroupConfigTaskInventoryView NfvInstGroupConfigTask
type NfvInstGroupConfigTaskInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	NfvInstGroupUuid string `json:"nfvInstGroupUuid,omitempty"`
	ConfigVersion int `json:"configVersion,omitempty"`
	ServiceUuid string `json:"serviceUuid,omitempty"`
	TaskName string `json:"taskName,omitempty"`
	TaskData string `json:"taskData,omitempty"`
	Path string `json:"path,omitempty"`
	CheckStatus bool `json:"checkStatus,omitempty"`
}

