// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSDingTalkAtPersonInventoryView SNSDingTalkAtPerson
type SNSDingTalkAtPersonInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Remark string `json:"remark,omitempty"`
}

