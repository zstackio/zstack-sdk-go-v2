// Copyright (c) ZStack.io, Inc.

package param

// IsVfNicAvailableInL3NetworkDetailParam IsVfNicAvailableInL3Network详细参数
type IsVfNicAvailableInL3NetworkDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"hostUuid" validate:"required"` // 必填
}

// IsVfNicAvailableInL3NetworkParam IsVfNicAvailableInL3Network请求参数
type IsVfNicAvailableInL3NetworkParam struct {
	BaseParam
	Params IsVfNicAvailableInL3NetworkDetailParam `json:"params"` // 详细参数
}

