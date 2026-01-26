// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectAccountRefInventoryView IAM2ProjectAccountRef
type IAM2ProjectAccountRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ProjectUuid string `json:"projectUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// QueryIAM2ProjectAccountRefView QueryIAM2ProjectAccountRef
type QueryIAM2ProjectAccountRefView struct {
	Inventories []IAM2ProjectAccountRefInventoryView `json:"inventories,omitempty"`
}

