// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateVmNicsForPortMirrorDetailParam GetCandidateVmNicsForPortMirror detail param
type GetCandidateVmNicsForPortMirrorDetailParam struct {
	PortMirrorUuid string `json:"portMirrorUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
}

// GetCandidateVmNicsForPortMirrorParam GetCandidateVmNicsForPortMirror request param
type GetCandidateVmNicsForPortMirrorParam struct {
	BaseParam
	Params GetCandidateVmNicsForPortMirrorDetailParam `json:"params"`
}
