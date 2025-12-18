// Copyright (c) ZStack.io, Inc.

package param

// ChangeSlbGroupDeployTypeDetailParam ChangeSlbGroupDeployType detail param
type ChangeSlbGroupDeployTypeDetailParam struct {
	SlbGroupUuid string `json:"slbGroupUuid" validate:"required"`
	DeployType string `json:"deployType" validate:"required"`
}

// ChangeSlbGroupDeployTypeParam ChangeSlbGroupDeployType request param
type ChangeSlbGroupDeployTypeParam struct {
	BaseParam
	Params ChangeSlbGroupDeployTypeDetailParam `json:"params"`
}
