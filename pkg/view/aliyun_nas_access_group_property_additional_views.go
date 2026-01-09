// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasAccessGroupPropertyView AliyunNasAccessGroupProperty
type AliyunNasAccessGroupPropertyView struct {
	RuleCount int `json:"ruleCount,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	NetworkType *string `json:"networkType,omitempty"`
}

