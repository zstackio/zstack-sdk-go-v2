// Copyright (c) ZStack.io, Inc.

package param

// GetVpcIPsecLogDetailParam GetVpcIPsecLog detail param
type GetVpcIPsecLogDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Lines int `json:"lines,omitempty"`
}

// GetVpcIPsecLogParam GetVpcIPsecLog request param
type GetVpcIPsecLogParam struct {
	BaseParam
	Params GetVpcIPsecLogDetailParam `json:"params"`
}
