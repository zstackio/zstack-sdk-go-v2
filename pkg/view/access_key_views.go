// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AccessKeyInventoryView AccessKey
type AccessKeyInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	UserUuid string `json:"userUuid,omitempty"`
	AccessKeyID string `json:"AccessKeyID,omitempty"`
	AccessKeySecret string `json:"AccessKeySecret,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

