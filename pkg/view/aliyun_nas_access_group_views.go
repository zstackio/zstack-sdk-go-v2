// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunNasAccessGroupInventoryView AliyunNasAccessGroup
type AliyunNasAccessGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest []AliyunNasAccessRuleInventoryView `json:"rules,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

