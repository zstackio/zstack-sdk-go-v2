// Copyright (c) ZStack.io, Inc.

package param

// SyncChronyServersDetailParam SyncChronyServers detail param
type SyncChronyServersDetailParam struct {
}

// SyncChronyServersParam SyncChronyServers request param
type SyncChronyServersParam struct {
	BaseParam
	Params SyncChronyServersDetailParam `json:"params"`
}
