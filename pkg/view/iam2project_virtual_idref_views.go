// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectVirtualIDRefInventoryView IAM2ProjectVirtualIDRef
type IAM2ProjectVirtualIDRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ProjectUuid *string `json:"projectUuid,omitempty"`
	VirtualIDUuid *string `json:"virtualIDUuid,omitempty"`
}

