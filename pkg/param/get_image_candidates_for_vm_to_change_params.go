// Copyright (c) ZStack.io, Inc.

package param

// GetImageCandidatesForVmToChangeDetailParam GetImageCandidatesForVmToChange detail param
type GetImageCandidatesForVmToChangeDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
}

// GetImageCandidatesForVmToChangeParam GetImageCandidatesForVmToChange request param
type GetImageCandidatesForVmToChangeParam struct {
	BaseParam
	Params GetImageCandidatesForVmToChangeDetailParam `json:"params"`
}
