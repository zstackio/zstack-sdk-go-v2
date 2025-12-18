// Copyright (c) ZStack.io, Inc.

package param

// MigrateVmDetailParam MigrateVm详细参数
type MigrateVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"hostUuid,omitempty"`
	rest bool `json:"migrateFromDestination,omitempty"`
	rest bool `json:"allowUnknown,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest int `json:"downTime,omitempty"`
}

// MigrateVmParam MigrateVm请求参数
type MigrateVmParam struct {
	BaseParam
	Params MigrateVmDetailParam `json:"params"` // 详细参数
}

