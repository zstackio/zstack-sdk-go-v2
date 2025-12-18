// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunNasAccessGroupRuleDetailParam CreateAliyunNasAccessGroupRule detail param
type CreateAliyunNasAccessGroupRuleDetailParam struct {
	AccessGroupUuid string `json:"accessGroupUuid" validate:"required"`
	SourceCidrIp string `json:"sourceCidrIp" validate:"required"`
	RwAccessType string `json:"rwAccessType,omitempty"`
	Priority int `json:"priority,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunNasAccessGroupRuleParam CreateAliyunNasAccessGroupRule request param
type CreateAliyunNasAccessGroupRuleParam struct {
	BaseParam
	Params CreateAliyunNasAccessGroupRuleDetailParam `json:"params"`
}
