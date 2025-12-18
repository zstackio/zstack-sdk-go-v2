// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunNasAccessGroupDetailParam DeleteAliyunNasAccessGroup detail param
type DeleteAliyunNasAccessGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunNasAccessGroupParam DeleteAliyunNasAccessGroup request param
type DeleteAliyunNasAccessGroupParam struct {
	BaseParam
	Params DeleteAliyunNasAccessGroupDetailParam `json:"params"`
}
