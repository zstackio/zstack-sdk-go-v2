// Copyright (c) ZStack.io, Inc.

package param

// EnableCdpTaskDetailParam EnableCdpTask详细参数
type EnableCdpTaskDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// EnableCdpTaskParam EnableCdpTask请求参数
type EnableCdpTaskParam struct {
	BaseParam
	Params EnableCdpTaskDetailParam `json:"params"` // 详细参数
}

