// Copyright (c) ZStack.io, Inc.

package param

// GetImageCandidatesForVmToChangeDetailParam GetImageCandidatesForVmToChange详细参数
type GetImageCandidatesForVmToChangeDetailParam struct {
	rest string `json:"vmInstanceUuid,omitempty"`
}

// GetImageCandidatesForVmToChangeParam GetImageCandidatesForVmToChange请求参数
type GetImageCandidatesForVmToChangeParam struct {
	BaseParam
	Params GetImageCandidatesForVmToChangeDetailParam `json:"params"` // 详细参数
}

