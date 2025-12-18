// Copyright (c) ZStack.io, Inc.

package param

// AddSNSWeComAtPersonDetailParam AddSNSWeComAtPerson详细参数
type AddSNSWeComAtPersonDetailParam struct {
	rest string `json:"userId" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"remark,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSNSWeComAtPersonParam AddSNSWeComAtPerson请求参数
type AddSNSWeComAtPersonParam struct {
	BaseParam
	Params AddSNSWeComAtPersonDetailParam `json:"params"` // 详细参数
}

