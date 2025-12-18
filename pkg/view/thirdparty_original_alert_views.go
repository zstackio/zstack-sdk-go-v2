// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ThirdpartyOriginalAlertInventoryView ThirdpartyOriginalAlert
type ThirdpartyOriginalAlertInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"thirdpartyPlatformUuid,omitempty"`
	rest string `json:"product,omitempty"`
	rest string `json:"service,omitempty"`
	rest string `json:"metric,omitempty"`
	rest string `json:"alertLevel,omitempty"`
	rest time.Time `json:"alertTime,omitempty"`
	rest string `json:"dimensions,omitempty"`
	rest string `json:"message,omitempty"`
	rest string `json:"dataSource,omitempty"`
	rest string `json:"sourceText,omitempty"`
	rest string `json:"readStatus,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
}

