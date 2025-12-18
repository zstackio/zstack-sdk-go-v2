// Copyright (c) ZStack.io, Inc.

package param

// DeleteAccessControlListDetailParam DeleteAccessControlList detail param
type DeleteAccessControlListDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAccessControlListParam DeleteAccessControlList request param
type DeleteAccessControlListParam struct {
	BaseParam
	Params DeleteAccessControlListDetailParam `json:"params"`
}
