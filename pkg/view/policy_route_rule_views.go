// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PolicyRouteRuleInventoryView PolicyRouteRule
type PolicyRouteRuleInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest int `json:"ruleNumber,omitempty"`
	rest string `json:"ruleSetUuid,omitempty"`
	rest string `json:"tableUuid,omitempty"`
	rest string `json:"destIp,omitempty"`
	rest string `json:"sourceIp,omitempty"`
	rest string `json:"destPort,omitempty"`
	rest string `json:"sourcePort,omitempty"`
	rest string `json:"protocol,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

