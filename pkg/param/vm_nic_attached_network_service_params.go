// Copyright (c) ZStack.io, Inc.

package param

// GetVmNicAttachedNetworkServiceDetailParam GetVmNicAttachedNetworkService详细参数
type GetVmNicAttachedNetworkServiceDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
}

// GetVmNicAttachedNetworkServiceParam GetVmNicAttachedNetworkService请求参数
type GetVmNicAttachedNetworkServiceParam struct {
	BaseParam
	Params GetVmNicAttachedNetworkServiceDetailParam `json:"params"` // 详细参数
}

