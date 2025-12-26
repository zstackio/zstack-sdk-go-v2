// Copyright (c) ZStack.io, Inc.

package view

// SyncAINginxConfigurationView SyncAINginxConfiguration
type SyncAINginxConfigurationView struct {
	UnSyncedRules []NginxRedirectRuleView `json:"unSyncedRules,omitempty"`
	Success bool `json:"success,omitempty"`
}

