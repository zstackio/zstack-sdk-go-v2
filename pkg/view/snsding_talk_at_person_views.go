// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSDingTalkAtPersonInventoryView SNSDingTalkAtPerson
type SNSDingTalkAtPersonInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"phoneNumber,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"remark,omitempty"`
}

