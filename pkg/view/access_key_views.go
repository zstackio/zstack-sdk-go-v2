// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AccessKeyInventoryView AccessKey
type AccessKeyInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"userUuid,omitempty"`
	rest string `json:"AccessKeyID,omitempty"`
	rest string `json:"AccessKeySecret,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

