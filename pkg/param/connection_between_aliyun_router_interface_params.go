// Copyright (c) ZStack.io, Inc.

package param

// StartConnectionBetweenAliyunRouterInterfaceDetailParam StartConnectionBetweenAliyunRouterInterface详细参数
type StartConnectionBetweenAliyunRouterInterfaceDetailParam struct {
	rest string `json:"vrouterInterfaceUuid" validate:"required"` // 必填
	rest string `json:"vbrInterfaceUuid" validate:"required"` // 必填
}

// StartConnectionBetweenAliyunRouterInterfaceParam StartConnectionBetweenAliyunRouterInterface请求参数
type StartConnectionBetweenAliyunRouterInterfaceParam struct {
	BaseParam
	Params StartConnectionBetweenAliyunRouterInterfaceDetailParam `json:"params"` // 详细参数
}

