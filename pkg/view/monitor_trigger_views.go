// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MonitorTriggerInventoryView MonitorTrigger
type MonitorTriggerInventoryView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Expression string `json:"expression,omitempty"`
	RecoveryExpression string `json:"recoveryExpression,omitempty"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
	Duration int `json:"duration,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid,omitempty"`
	LastStatusChangeTime time.Time `json:"lastStatusChangeTime,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

