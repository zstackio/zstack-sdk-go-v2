// Copyright (c) ZStack.io, Inc.

package param

// KvmRunShellDetailParam KvmRunShell detail param
type KvmRunShellDetailParam struct {
	HostUuids []string `json:"hostUuids" validate:"required"`
	Script string `json:"script" validate:"required"`
}

// KvmRunShellParam KvmRunShell request param
type KvmRunShellParam struct {
	BaseParam
	Params KvmRunShellDetailParam `json:"params"`
}
