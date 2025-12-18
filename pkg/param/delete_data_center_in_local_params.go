// Copyright (c) ZStack.io, Inc.

package param

// DeleteDataCenterInLocalDetailParam DeleteDataCenterInLocal detail param
type DeleteDataCenterInLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteDataCenterInLocalParam DeleteDataCenterInLocal request param
type DeleteDataCenterInLocalParam struct {
	BaseParam
	Params DeleteDataCenterInLocalDetailParam `json:"params"`
}
