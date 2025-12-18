// Copyright (c) ZStack.io, Inc.

package param

// CheckIAM2VirtualIDConfigFileDetailParam CheckIAM2VirtualIDConfigFile detail param
type CheckIAM2VirtualIDConfigFileDetailParam struct {
	VirtualIDInfos string `json:"virtualIDInfos" validate:"required"`
}

// CheckIAM2VirtualIDConfigFileParam CheckIAM2VirtualIDConfigFile request param
type CheckIAM2VirtualIDConfigFileParam struct {
	BaseParam
	Params CheckIAM2VirtualIDConfigFileDetailParam `json:"params"`
}
