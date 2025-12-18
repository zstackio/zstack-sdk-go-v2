// Copyright (c) ZStack.io, Inc.

package param

// DeleteCdpTaskDataDetailParam DeleteCdpTaskData detail param
type DeleteCdpTaskDataDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteCdpTaskDataParam DeleteCdpTaskData request param
type DeleteCdpTaskDataParam struct {
	BaseParam
	Params DeleteCdpTaskDataDetailParam `json:"params"`
}
