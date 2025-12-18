// Copyright (c) ZStack.io, Inc.

package param

// IsReadyToGoDetailParam IsReadyToGo详细参数
type IsReadyToGoDetailParam struct {
	rest string `json:"managementNodeId,omitempty"`
}

// IsReadyToGoParam IsReadyToGo请求参数
type IsReadyToGoParam struct {
	BaseParam
	Params IsReadyToGoDetailParam `json:"params"` // 详细参数
}

