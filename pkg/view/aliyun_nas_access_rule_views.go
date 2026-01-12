// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasAccessRuleInventoryView AliyunNasAccessRule
type AliyunNasAccessRuleInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccessGroupUuid *string `json:"accessGroupUuid,omitempty"`
	SourceCidr *string `json:"sourceCidr,omitempty"`
	Rule *string `json:"rule,omitempty"`
	Priority *int `json:"priority,omitempty"`
	UserAccess *string `json:"userAccess,omitempty"`
	RuleId *string `json:"ruleId,omitempty"`
}

// CreateAliyunNasAccessGroupRuleEventView CreateAliyunNasAccessGroupRuleEvent
type CreateAliyunNasAccessGroupRuleEventView struct {
	Inventory AliyunNasAccessRuleInventoryView `json:"inventory,omitempty"`
}

