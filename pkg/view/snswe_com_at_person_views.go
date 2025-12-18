// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSWeComAtPersonInventoryView SNSWeComAtPerson
type SNSWeComAtPersonInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"userId,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"remark,omitempty"`
}

