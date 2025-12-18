// Copyright (c) ZStack.io, Inc.

package param

// SyncAINginxConfigurationDetailParam SyncAINginxConfiguration detail param
type SyncAINginxConfigurationDetailParam struct {
	GroupUuids []string `json:"groupUuids,omitempty"`
	DryRun bool `json:"dryRun,omitempty"`
	SyncAll bool `json:"syncAll,omitempty"`
}

// SyncAINginxConfigurationParam SyncAINginxConfiguration request param
type SyncAINginxConfigurationParam struct {
	BaseParam
	Params SyncAINginxConfigurationDetailParam `json:"params"`
}
