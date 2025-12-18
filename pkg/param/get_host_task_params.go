// Copyright (c) ZStack.io, Inc.

package param

// GetHostTaskDetailParam GetHostTask detail param
type GetHostTaskDetailParam struct {
	HostUuids []string `json:"hostUuids" validate:"required"`
	SyncSignatures []string `json:"syncSignatures,omitempty"`
}

// GetHostTaskParam GetHostTask request param
type GetHostTaskParam struct {
	BaseParam
	Params GetHostTaskDetailParam `json:"params"`
}
