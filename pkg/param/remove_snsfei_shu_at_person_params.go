// Copyright (c) ZStack.io, Inc.

package param

// RemoveSNSFeiShuAtPersonDetailParam RemoveSNSFeiShuAtPerson详细参数
type RemoveSNSFeiShuAtPersonDetailParam struct {
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"userId" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveSNSFeiShuAtPersonParam RemoveSNSFeiShuAtPerson请求参数
type RemoveSNSFeiShuAtPersonParam struct {
	BaseParam
	Params RemoveSNSFeiShuAtPersonDetailParam `json:"params"` // 详细参数
}

