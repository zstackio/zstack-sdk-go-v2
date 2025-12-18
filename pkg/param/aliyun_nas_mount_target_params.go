// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunNasMountTargetDetailParam CreateAliyunNasMountTarget详细参数
type CreateAliyunNasMountTargetDetailParam struct {
	rest string `json:"nasAccessGroupUuid" validate:"required"` // 必填
	rest string `json:"vSwitchUuid,omitempty"`
	rest string `json:"nasFSUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateAliyunNasMountTargetParam CreateAliyunNasMountTarget请求参数
type CreateAliyunNasMountTargetParam struct {
	BaseParam
	Params CreateAliyunNasMountTargetDetailParam `json:"params"` // 详细参数
}

