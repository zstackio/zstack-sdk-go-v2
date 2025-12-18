// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunNasAccessGroupDetailParam AddAliyunNasAccessGroup详细参数
type AddAliyunNasAccessGroupDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"groupName" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasAccessGroupParam AddAliyunNasAccessGroup请求参数
type AddAliyunNasAccessGroupParam struct {
	BaseParam
	Params AddAliyunNasAccessGroupDetailParam `json:"params"` // 详细参数
}

