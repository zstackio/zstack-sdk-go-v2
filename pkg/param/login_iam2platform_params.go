// Copyright (c) ZStack.io, Inc.

package param

// LoginIAM2PlatformDetailParam LoginIAM2Platform detail param
type LoginIAM2PlatformDetailParam struct {
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2PlatformParam LoginIAM2Platform request param
type LoginIAM2PlatformParam struct {
	BaseParam
	Params LoginIAM2PlatformDetailParam `json:"params"`
}
