// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSTextTemplateInventoryView SNSTextTemplate
type SNSTextTemplateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"applicationPlatformType,omitempty"`
	rest string `json:"subject,omitempty"`
	rest string `json:"recoverySubject,omitempty"`
	rest string `json:"template,omitempty"`
	rest string `json:"recoveryTemplate,omitempty"`
	rest bool `json:"defaultTemplate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"type,omitempty"`
}

