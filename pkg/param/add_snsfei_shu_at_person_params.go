// Copyright (c) ZStack.io, Inc.

package param

// AddSNSFeiShuAtPersonDetailParam AddSNSFeiShuAtPerson详细参数
type AddSNSFeiShuAtPersonDetailParam struct {
	rest string `json:"userId" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"remark,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSNSFeiShuAtPersonParam AddSNSFeiShuAtPerson请求参数
type AddSNSFeiShuAtPersonParam struct {
	BaseParam
	Params AddSNSFeiShuAtPersonDetailParam `json:"params"` // 详细参数
}

