// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TemplateCustomParamInventoryView TemplateCustomParam
type TemplateCustomParamInventoryView struct {
	TemplateUuid string `json:"templateUuid,omitempty"`
	Param string `json:"param,omitempty"`
}

