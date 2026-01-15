// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateEcsVpcParamDetail UpdateEcsVpc detail param
type UpdateEcsVpcParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateEcsVpcParam UpdateEcsVpc request param
type UpdateEcsVpcParam struct {
	BaseParam
	UpdateEcsVpc UpdateEcsVpcParamDetail `json:"updateEcsVpc"`
}
