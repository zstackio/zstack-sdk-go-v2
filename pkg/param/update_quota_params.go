// Copyright (c) ZStack.io, Inc.

package param

// UpdateQuotaDetailParam UpdateQuota detail param
type UpdateQuotaDetailParam struct {
	IdentityUuid string `json:"identityUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Value int64 `json:"value" validate:"required"`
}

// UpdateQuotaParam UpdateQuota request param
type UpdateQuotaParam struct {
	BaseParam
	Params UpdateQuotaDetailParam `json:"params"`
}
