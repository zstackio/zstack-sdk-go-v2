// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunNasAccessGroupDetailParam UpdateAliyunNasAccessGroup detail param
type UpdateAliyunNasAccessGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
}

// UpdateAliyunNasAccessGroupParam UpdateAliyunNasAccessGroup request param
type UpdateAliyunNasAccessGroupParam struct {
	BaseParam
	Params UpdateAliyunNasAccessGroupDetailParam `json:"params"`
}
