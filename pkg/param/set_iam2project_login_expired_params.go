// Copyright (c) ZStack.io, Inc.

package param

// SetIAM2ProjectLoginExpiredDetailParam SetIAM2ProjectLoginExpired detail param
type SetIAM2ProjectLoginExpiredDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	LoginExpired string `json:"loginExpired" validate:"required"`
	LoginExpiredAttributeUuid string `json:"loginExpiredAttributeUuid,omitempty"`
}

// SetIAM2ProjectLoginExpiredParam SetIAM2ProjectLoginExpired request param
type SetIAM2ProjectLoginExpiredParam struct {
	BaseParam
	Params SetIAM2ProjectLoginExpiredDetailParam `json:"params"`
}
