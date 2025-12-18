// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2ProjectLoginExpiredDetailParam RemoveIAM2ProjectLoginExpired detail param
type RemoveIAM2ProjectLoginExpiredDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	AttributeUuid string `json:"attributeUuid" validate:"required"`
}

// RemoveIAM2ProjectLoginExpiredParam RemoveIAM2ProjectLoginExpired request param
type RemoveIAM2ProjectLoginExpiredParam struct {
	BaseParam
	Params RemoveIAM2ProjectLoginExpiredDetailParam `json:"params"`
}
