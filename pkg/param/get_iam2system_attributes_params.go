// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2SystemAttributesDetailParam GetIAM2SystemAttributes detail param
type GetIAM2SystemAttributesDetailParam struct {
}

// GetIAM2SystemAttributesParam GetIAM2SystemAttributes request param
type GetIAM2SystemAttributesParam struct {
	BaseParam
	Params GetIAM2SystemAttributesDetailParam `json:"params"`
}
