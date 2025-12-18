// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunNasAccessRuleInventoryView AliyunNasAccessRule
type AliyunNasAccessRuleInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"accessGroupUuid,omitempty"`
	rest string `json:"sourceCidr,omitempty"`
	rest string `json:"rule,omitempty"`
	rest int `json:"priority,omitempty"`
	rest string `json:"userAccess,omitempty"`
	rest string `json:"ruleId,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

