// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsInstanceTypeView EcsInstanceType
type EcsInstanceTypeView struct {
	TypeId     string `json:"typeId,omitempty"`
	Cpu        int    `json:"cpu,omitempty"`
	Memory     int64  `json:"memory,omitempty"`
	TypeFamily string `json:"typeFamily,omitempty"`
	Generation string `json:"generation,omitempty"`
}
