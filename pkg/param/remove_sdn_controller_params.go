// Copyright (c) ZStack.io, Inc.

package param

// RemoveSdnControllerDetailParam RemoveSdnController detail param
type RemoveSdnControllerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveSdnControllerParam RemoveSdnController request param
type RemoveSdnControllerParam struct {
	BaseParam
	Params RemoveSdnControllerDetailParam `json:"params"`
}
