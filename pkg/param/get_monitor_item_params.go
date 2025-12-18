// Copyright (c) ZStack.io, Inc.

package param

// GetMonitorItemDetailParam GetMonitorItem detail param
type GetMonitorItemDetailParam struct {
	ResourceType string `json:"resourceType" validate:"required"`
}

// GetMonitorItemParam GetMonitorItem request param
type GetMonitorItemParam struct {
	BaseParam
	Params GetMonitorItemDetailParam `json:"params"`
}
