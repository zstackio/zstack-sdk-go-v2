// Copyright (c) ZStack.io, Inc.

package param

// ChangeSdnControllerDetailParam ChangeSdnController detail param
type ChangeSdnControllerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
	VlanRanges []string `json:"vlanRanges,omitempty"`
}

// ChangeSdnControllerParam ChangeSdnController request param
type ChangeSdnControllerParam struct {
	BaseParam
	Params ChangeSdnControllerDetailParam `json:"params"`
}
