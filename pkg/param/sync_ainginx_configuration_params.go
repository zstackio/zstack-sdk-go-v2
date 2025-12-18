// Copyright (c) ZStack.io, Inc.

package param

// SyncAINginxConfigurationDetailParam SyncAINginxConfiguration详细参数
type SyncAINginxConfigurationDetailParam struct {
	rest []string `json:"groupUuids,omitempty"`
	rest bool `json:"dryRun,omitempty"`
	rest bool `json:"syncAll,omitempty"`
}

// SyncAINginxConfigurationParam SyncAINginxConfiguration请求参数
type SyncAINginxConfigurationParam struct {
	BaseParam
	Params SyncAINginxConfigurationDetailParam `json:"params"` // 详细参数
}

