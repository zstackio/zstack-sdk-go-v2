// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectVirtualIDGroupRefInventoryView IAM2ProjectVirtualIDGroupRef
type IAM2ProjectVirtualIDGroupRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ProjectUuid *string `json:"projectUuid,omitempty"`
	GroupUuid *string `json:"groupUuid,omitempty"`
}

