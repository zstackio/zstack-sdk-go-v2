// Copyright (c) ZStack.io, Inc.

package param

// GetModelCenterServicesDetailParam GetModelCenterServices detail param
type GetModelCenterServicesDetailParam struct {
	ModelCenterUuids []string `json:"modelCenterUuids,omitempty"`
}

// GetModelCenterServicesParam GetModelCenterServices request param
type GetModelCenterServicesParam struct {
	BaseParam
	Params GetModelCenterServicesDetailParam `json:"params"`
}
