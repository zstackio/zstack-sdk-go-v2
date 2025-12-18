// Copyright (c) ZStack.io, Inc.

package param

// UpdateOrganizationQuotaDetailParam UpdateOrganizationQuota detail param
type UpdateOrganizationQuotaDetailParam struct {
	IdentityUuid string `json:"identityUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Value int64 `json:"value" validate:"required"`
}

// UpdateOrganizationQuotaParam UpdateOrganizationQuota request param
type UpdateOrganizationQuotaParam struct {
	BaseParam
	Params UpdateOrganizationQuotaDetailParam `json:"params"`
}
