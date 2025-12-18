// Copyright (c) ZStack.io, Inc.

package param

// EnableCdpTaskDetailParam EnableCdpTask detail param
type EnableCdpTaskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// EnableCdpTaskParam EnableCdpTask request param
type EnableCdpTaskParam struct {
	BaseParam
	Params EnableCdpTaskDetailParam `json:"params"`
}
