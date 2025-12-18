// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SnmpAgentInventoryView SnmpAgent
type SnmpAgentInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"readCommunity,omitempty"`
	rest string `json:"userName,omitempty"`
	rest string `json:"authAlgorithm,omitempty"`
	rest string `json:"authPassword,omitempty"`
	rest string `json:"privacyAlgorithm,omitempty"`
	rest string `json:"privacyPassword,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"securityLevel,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

