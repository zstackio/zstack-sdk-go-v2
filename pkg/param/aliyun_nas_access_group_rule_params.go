// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunNasAccessGroupRuleDetailParam CreateAliyunNasAccessGroupRule详细参数
type CreateAliyunNasAccessGroupRuleDetailParam struct {
	rest string `json:"accessGroupUuid" validate:"required"` // 必填
	rest string `json:"sourceCidrIp" validate:"required"` // 必填
	rest string `json:"rwAccessType,omitempty"`
	rest int `json:"priority,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateAliyunNasAccessGroupRuleParam CreateAliyunNasAccessGroupRule请求参数
type CreateAliyunNasAccessGroupRuleParam struct {
	BaseParam
	Params CreateAliyunNasAccessGroupRuleDetailParam `json:"params"` // 详细参数
}

