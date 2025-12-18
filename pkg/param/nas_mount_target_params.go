// Copyright (c) ZStack.io, Inc.

package param

// QueryNasMountTargetDetailParam QueryNasMountTarget详细参数
type QueryNasMountTargetDetailParam struct {
	rest []interface{} `json:"conditions" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
	rest bool `json:"count,omitempty"`
	rest string `json:"groupBy,omitempty"`
	rest bool `json:"replyWithCount,omitempty"`
	rest string `json:"filterName,omitempty"`
	rest string `json:"sortBy,omitempty"`
	rest string `json:"sortDirection,omitempty"`
	rest []string `json:"fields,omitempty"`
}

// QueryNasMountTargetParam QueryNasMountTarget请求参数
type QueryNasMountTargetParam struct {
	BaseParam
	Params QueryNasMountTargetDetailParam `json:"params"` // 详细参数
}

