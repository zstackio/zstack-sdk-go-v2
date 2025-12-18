// Copyright (c) ZStack.io, Inc.

package param

// IsReadyToGoDetailParam IsReadyToGo detail param
type IsReadyToGoDetailParam struct {
	ManagementNodeId string `json:"managementNodeId,omitempty"`
}

// IsReadyToGoParam IsReadyToGo request param
type IsReadyToGoParam struct {
	BaseParam
	Params IsReadyToGoDetailParam `json:"params"`
}
