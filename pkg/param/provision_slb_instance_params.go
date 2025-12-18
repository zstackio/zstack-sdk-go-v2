// Copyright (c) ZStack.io, Inc.

package param

// ProvisionSlbInstanceDetailParam ProvisionSlbInstance detail param
type ProvisionSlbInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ProvisionSlbInstanceParam ProvisionSlbInstance request param
type ProvisionSlbInstanceParam struct {
	BaseParam
	Params ProvisionSlbInstanceDetailParam `json:"params"`
}
