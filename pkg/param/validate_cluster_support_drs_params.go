// Copyright (c) ZStack.io, Inc.

package param

// ValidateClusterSupportDRSDetailParam ValidateClusterSupportDRS详细参数
type ValidateClusterSupportDRSDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// ValidateClusterSupportDRSParam ValidateClusterSupportDRS请求参数
type ValidateClusterSupportDRSParam struct {
	BaseParam
	Params ValidateClusterSupportDRSDetailParam `json:"params"` // 详细参数
}

