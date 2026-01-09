// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AutoScalingGroupInstanceInventoryView AutoScalingGroupInstance
type AutoScalingGroupInstanceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	InstanceUuid *string `json:"instanceUuid,omitempty"`
	ScalingGroupUuid *string `json:"scalingGroupUuid,omitempty"`
	TemplateUuid *string `json:"templateUuid,omitempty"`
	ScalingGroupActivityUuid *string `json:"scalingGroupActivityUuid,omitempty"`
	Status *string `json:"status,omitempty"`
	HealthStatus *string `json:"healthStatus,omitempty"`
	Description *string `json:"description,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	ProtectionStrategy *string `json:"protectionStrategy,omitempty"`
}

// UpdateAutoScalingGroupInstanceEventView UpdateAutoScalingGroupInstanceEvent
type UpdateAutoScalingGroupInstanceEventView struct {
	Inventory AutoScalingGroupInstanceInventoryView `json:"inventory,omitempty"`
}

// DeleteAutoScalingGroupInstanceEventView DeleteAutoScalingGroupInstanceEvent
type DeleteAutoScalingGroupInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryAutoScalingGroupInstanceView QueryAutoScalingGroupInstance
type QueryAutoScalingGroupInstanceView struct {
	Inventories []AutoScalingGroupInstanceInventoryView `json:"inventories,omitempty"`
}

