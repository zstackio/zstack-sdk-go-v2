// Copyright (c) ZStack.io, Inc.

package param

// UpdateOrganizationQuotaDetailParam UpdateOrganizationQuota详细参数
type UpdateOrganizationQuotaDetailParam struct {
	rest string `json:"identityUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest int64 `json:"value" validate:"required"` // 必填
}

// UpdateOrganizationQuotaParam UpdateOrganizationQuota请求参数
type UpdateOrganizationQuotaParam struct {
	BaseParam
	Params UpdateOrganizationQuotaDetailParam `json:"params"` // 详细参数
}

