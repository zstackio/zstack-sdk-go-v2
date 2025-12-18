// Copyright (c) ZStack.io, Inc.

package param

// ValidateClusterSupportDRSDetailParam ValidateClusterSupportDRS detail param
type ValidateClusterSupportDRSDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// ValidateClusterSupportDRSParam ValidateClusterSupportDRS request param
type ValidateClusterSupportDRSParam struct {
	BaseParam
	Params ValidateClusterSupportDRSDetailParam `json:"params"`
}
