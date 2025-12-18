// Copyright (c) ZStack.io, Inc.

package param

// KvmRunShellDetailParam KvmRunShell详细参数
type KvmRunShellDetailParam struct {
	rest []string `json:"hostUuids" validate:"required"` // 必填
	rest string `json:"script" validate:"required"` // 必填
}

// KvmRunShellParam KvmRunShell请求参数
type KvmRunShellParam struct {
	BaseParam
	Params KvmRunShellDetailParam `json:"params"` // 详细参数
}

