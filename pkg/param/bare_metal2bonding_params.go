// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateBareMetal2BondingParamDetail CreateBareMetal2Bonding detail param
type CreateBareMetal2BondingParamDetail struct {
	ChassisUuid string `json:"chassisUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Mode int `json:"mode" validate:"required"`
	Slaves string `json:"slaves" validate:"required"`
	Opts *string `json:"opts,omitempty"`
}

// CreateBareMetal2BondingParam CreateBareMetal2Bonding request param
type CreateBareMetal2BondingParam struct {
	BaseParam
	Params CreateBareMetal2BondingParamDetail `json:"params"`
}
