// Copyright (c) ZStack.io, Inc.

package param

// DeleteIdentityZoneInLocalDetailParam DeleteIdentityZoneInLocal detail param
type DeleteIdentityZoneInLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIdentityZoneInLocalParam DeleteIdentityZoneInLocal request param
type DeleteIdentityZoneInLocalParam struct {
	BaseParam
	Params DeleteIdentityZoneInLocalDetailParam `json:"params"`
}
