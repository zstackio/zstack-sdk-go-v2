// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmSchedulingRuleInventoryView VmSchedulingRule
type VmSchedulingRuleInventoryView struct {
	Rule string `json:"rule,omitempty"`
	Mode string `json:"mode,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Policy string `json:"policy,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
	Appliance string `json:"appliance,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Usages []AffinityGroupUsageInventoryView `json:"usages,omitempty"`
}

