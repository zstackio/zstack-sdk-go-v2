// Copyright (c) ZStack.io, Inc.

package param

// GetAliyunNasMountTargetRemoteDetailParam GetAliyunNasMountTargetRemote详细参数
type GetAliyunNasMountTargetRemoteDetailParam struct {
	rest string `json:"nasFSUuid" validate:"required"` // 必填
	rest string `json:"mountDomain,omitempty"`
}

// GetAliyunNasMountTargetRemoteParam GetAliyunNasMountTargetRemote请求参数
type GetAliyunNasMountTargetRemoteParam struct {
	BaseParam
	Params GetAliyunNasMountTargetRemoteDetailParam `json:"params"` // 详细参数
}

