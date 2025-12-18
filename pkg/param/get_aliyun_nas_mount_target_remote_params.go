// Copyright (c) ZStack.io, Inc.

package param

// GetAliyunNasMountTargetRemoteDetailParam GetAliyunNasMountTargetRemote detail param
type GetAliyunNasMountTargetRemoteDetailParam struct {
	NasFSUuid string `json:"nasFSUuid" validate:"required"`
	MountDomain string `json:"mountDomain,omitempty"`
}

// GetAliyunNasMountTargetRemoteParam GetAliyunNasMountTargetRemote request param
type GetAliyunNasMountTargetRemoteParam struct {
	BaseParam
	Params GetAliyunNasMountTargetRemoteDetailParam `json:"params"`
}
