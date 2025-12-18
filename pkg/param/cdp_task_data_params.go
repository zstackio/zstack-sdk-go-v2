// Copyright (c) ZStack.io, Inc.

package param

// DeleteCdpTaskDataDetailParam DeleteCdpTaskData详细参数
type DeleteCdpTaskDataDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// DeleteCdpTaskDataParam DeleteCdpTaskData请求参数
type DeleteCdpTaskDataParam struct {
	BaseParam
	Params DeleteCdpTaskDataDetailParam `json:"params"` // 详细参数
}

