// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SecurityGroupInventoryView SecurityGroup
type SecurityGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest string `json:"vSwitchType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest int64 `json:"internalId,omitempty"`
	rest []SecurityGroupRuleInventoryView `json:"rules,omitempty"`
	rest []string `json:"attachedL3NetworkUuids,omitempty"`
}

