// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TemplateView Template
type TemplateView struct {
	Attributes []AttributeView `json:"attributes,omitempty"`
	Quota map[string]int64 `json:"quota,omitempty"`
}

