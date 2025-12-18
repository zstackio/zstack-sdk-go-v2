// Copyright (c) ZStack.io, Inc.

package param

// GetEipAttachableVmNicsDetailParam GetEipAttachableVmNics详细参数
type GetEipAttachableVmNicsDetailParam struct {
	rest string `json:"eipUuid,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest string `json:"vmName,omitempty"`
	rest string `json:"networkServiceProvider,omitempty"`
	rest bool `json:"attachedToVm,omitempty"`
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetEipAttachableVmNicsParam GetEipAttachableVmNics请求参数
type GetEipAttachableVmNicsParam struct {
	BaseParam
	Params GetEipAttachableVmNicsDetailParam `json:"params"` // 详细参数
}

