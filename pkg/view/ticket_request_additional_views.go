// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketRequestView TicketRequest
type TicketRequestView struct {
	RequestName *string `json:"requestName,omitempty"`
	ApiName *string `json:"apiName,omitempty"`
	ExecuteTimes int `json:"executeTimes,omitempty"`
	ApiBody interface{} `json:"apiBody,omitempty"`
}

