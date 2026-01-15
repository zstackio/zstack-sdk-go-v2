// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ContainerImageTagInventoryView ContainerImageTag
type ContainerImageTagInventoryView struct {
	BaseInfoView
	BaseTimeView
	Size int64 `json:"size,omitempty"`
	SizeStr string `json:"sizeStr,omitempty"`
	Digest string `json:"digest,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	Address string `json:"address,omitempty"`
	PullCommand string `json:"pullCommand,omitempty"`
	Architectures string `json:"architectures,omitempty"`
}

