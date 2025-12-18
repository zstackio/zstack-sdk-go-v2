// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyRouteRuleInventoryView PolicyRouteRule
type PolicyRouteRuleInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	RuleNumber int `json:"ruleNumber,omitempty"`
	RuleSetUuid string `json:"ruleSetUuid,omitempty"`
	TableUuid string `json:"tableUuid,omitempty"`
	DestIp string `json:"destIp,omitempty"`
	SourceIp string `json:"sourceIp,omitempty"`
	DestPort string `json:"destPort,omitempty"`
	SourcePort string `json:"sourcePort,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

