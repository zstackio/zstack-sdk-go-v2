// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsSecurityGroupInLocalDetailParam DeleteEcsSecurityGroupInLocal detail param
type DeleteEcsSecurityGroupInLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsSecurityGroupInLocalParam DeleteEcsSecurityGroupInLocal request param
type DeleteEcsSecurityGroupInLocalParam struct {
	BaseParam
	Params DeleteEcsSecurityGroupInLocalDetailParam `json:"params"`
}
