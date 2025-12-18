// Copyright (c) ZStack.io, Inc.

package param

// UpdateVRouterOspfAreaDetailParam UpdateVRouterOspfArea detail param
type UpdateVRouterOspfAreaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	AreaAuth string `json:"areaAuth,omitempty"`
	AreaType string `json:"areaType,omitempty"`
	Password string `json:"password,omitempty"`
	KeyId int `json:"keyId,omitempty"`
}

// UpdateVRouterOspfAreaParam UpdateVRouterOspfArea request param
type UpdateVRouterOspfAreaParam struct {
	BaseParam
	Params UpdateVRouterOspfAreaDetailParam `json:"params"`
}
