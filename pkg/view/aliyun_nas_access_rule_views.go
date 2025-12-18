// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasAccessRuleInventoryView AliyunNasAccessRule
type AliyunNasAccessRuleInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccessGroupUuid string `json:"accessGroupUuid,omitempty"`
	SourceCidr string `json:"sourceCidr,omitempty"`
	Rule string `json:"rule,omitempty"`
	Priority int `json:"priority,omitempty"`
	UserAccess string `json:"userAccess,omitempty"`
	RuleId string `json:"ruleId,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

