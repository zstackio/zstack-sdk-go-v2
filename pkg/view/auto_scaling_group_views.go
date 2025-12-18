// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AutoScalingGroupInventoryView AutoScalingGroup
type AutoScalingGroupInventoryView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	ScalingResourceType string `json:"scalingResourceType,omitempty"`
	State string `json:"state,omitempty"`
	DefaultCooldown int64 `json:"defaultCooldown,omitempty"`
	Description string `json:"description,omitempty"`
	MinResourceSize int `json:"minResourceSize,omitempty"`
	MaxResourceSize int `json:"maxResourceSize,omitempty"`
	RemovalPolicy string `json:"removalPolicy,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	AttachedTemplates []string `json:"attachedTemplates,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
}

