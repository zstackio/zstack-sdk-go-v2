// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasAccessGroupInventoryView AliyunNasAccessGroup
type AliyunNasAccessGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	Rules []AliyunNasAccessRuleInventoryView `json:"rules,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

