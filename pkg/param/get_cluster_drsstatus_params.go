// Copyright (c) ZStack.io, Inc.

package param

// GetClusterDRSStatusDetailParam GetClusterDRSStatus detail param
type GetClusterDRSStatusDetailParam struct {
	DrsUuid string `json:"drsUuid" validate:"required"`
}

// GetClusterDRSStatusParam GetClusterDRSStatus request param
type GetClusterDRSStatusParam struct {
	BaseParam
	Params GetClusterDRSStatusDetailParam `json:"params"`
}
