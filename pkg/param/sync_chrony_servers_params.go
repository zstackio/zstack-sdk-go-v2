// Copyright (c) ZStack.io, Inc.

package param

// SyncChronyServersDetailParam SyncChronyServers详细参数
type SyncChronyServersDetailParam struct {
}

// SyncChronyServersParam SyncChronyServers请求参数
type SyncChronyServersParam struct {
	BaseParam
	Params SyncChronyServersDetailParam `json:"params"` // 详细参数
}

