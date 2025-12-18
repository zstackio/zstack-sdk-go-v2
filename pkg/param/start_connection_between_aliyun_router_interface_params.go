// Copyright (c) ZStack.io, Inc.

package param

// StartConnectionBetweenAliyunRouterInterfaceDetailParam StartConnectionBetweenAliyunRouterInterface detail param
type StartConnectionBetweenAliyunRouterInterfaceDetailParam struct {
	VrouterInterfaceUuid string `json:"vrouterInterfaceUuid" validate:"required"`
	VbrInterfaceUuid string `json:"vbrInterfaceUuid" validate:"required"`
}

// StartConnectionBetweenAliyunRouterInterfaceParam StartConnectionBetweenAliyunRouterInterface request param
type StartConnectionBetweenAliyunRouterInterfaceParam struct {
	BaseParam
	Params StartConnectionBetweenAliyunRouterInterfaceDetailParam `json:"params"`
}
