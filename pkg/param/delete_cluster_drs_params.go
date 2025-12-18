// Copyright (c) ZStack.io, Inc.

package param

// DeleteClusterDRSDetailParam DeleteClusterDRS detail param
type DeleteClusterDRSDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteClusterDRSParam DeleteClusterDRS request param
type DeleteClusterDRSParam struct {
	BaseParam
	Params DeleteClusterDRSDetailParam `json:"params"`
}
