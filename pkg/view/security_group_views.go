// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SecurityGroupInventoryView SecurityGroup
type SecurityGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	InternalId int64 `json:"internalId,omitempty"`
	Rules []SecurityGroupRuleInventoryView `json:"rules,omitempty"`
	AttachedL3NetworkUuids []string `json:"attachedL3NetworkUuids,omitempty"`
}

