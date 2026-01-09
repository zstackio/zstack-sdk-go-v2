// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ConsoleInventoryView Console
type ConsoleInventoryView struct {
	Scheme *string `json:"scheme,omitempty"`
	TargetScheme *string `json:"targetScheme,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Port int `json:"port,omitempty"`
	Token *string `json:"token,omitempty"`
	Version *string `json:"version,omitempty"`
	ExpiredDate *time.Time `json:"expiredDate,omitempty"`
}

