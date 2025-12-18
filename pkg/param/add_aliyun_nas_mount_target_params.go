// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunNasMountTargetDetailParam AddAliyunNasMountTarget详细参数
type AddAliyunNasMountTargetDetailParam struct {
	rest string `json:"nasFSUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"mountDomain" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasMountTargetParam AddAliyunNasMountTarget请求参数
type AddAliyunNasMountTargetParam struct {
	BaseParam
	Params AddAliyunNasMountTargetDetailParam `json:"params"` // 详细参数
}

