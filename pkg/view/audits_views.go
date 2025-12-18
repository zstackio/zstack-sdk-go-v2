// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AuditsInventoryView Audits
type AuditsInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest int64 `json:"createTime,omitempty"`
	rest string `json:"apiName,omitempty"`
	rest string `json:"clientBrowser,omitempty"`
	rest string `json:"clientIp,omitempty"`
	rest int64 `json:"duration,omitempty"`
	rest string `json:"operator,omitempty"`
	rest string `json:"requestDump,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"requestUuid,omitempty"`
	rest string `json:"operatorAccountUuid,omitempty"`
	rest string `json:"responseDump,omitempty"`
	rest bool `json:"success,omitempty"`
	rest string `json:"signedText,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"responseUuid,omitempty"`
	rest string `json:"resourceName,omitempty"`
	rest int64 `json:"startTime,omitempty"`
}

