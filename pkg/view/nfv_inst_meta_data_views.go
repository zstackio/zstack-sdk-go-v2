// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NfvInstMetaDataInventoryView NfvInstMetaData
type NfvInstMetaDataInventoryView struct {
	BaseInfoView
	BaseTimeView
	AgentVersion string `json:"agentVersion,omitempty"`
	NetOsVersion string `json:"netOsVersion,omitempty"`
	BaseOsVersion string `json:"baseOsVersion,omitempty"`
	KernelVersion string `json:"kernelVersion,omitempty"`
}

