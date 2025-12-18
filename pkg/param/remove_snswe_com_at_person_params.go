// Copyright (c) ZStack.io, Inc.

package param

// RemoveSNSWeComAtPersonDetailParam RemoveSNSWeComAtPerson详细参数
type RemoveSNSWeComAtPersonDetailParam struct {
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"userId" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveSNSWeComAtPersonParam RemoveSNSWeComAtPerson请求参数
type RemoveSNSWeComAtPersonParam struct {
	BaseParam
	Params RemoveSNSWeComAtPersonDetailParam `json:"params"` // 详细参数
}

