// Copyright (c) ZStack.io, Inc.

package param

// CreateVmNicDetailParam CreateVmNic详细参数
type CreateVmNicDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"ip,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVmNicParam CreateVmNic请求参数
type CreateVmNicParam struct {
	BaseParam
	Params CreateVmNicDetailParam `json:"params"` // 详细参数
}

