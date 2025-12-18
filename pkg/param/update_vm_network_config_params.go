// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmNetworkConfigDetailParam UpdateVmNetworkConfig detail param
type UpdateVmNetworkConfigDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	VmNicUuids []string `json:"vmNicUuids" validate:"required"`
}

// UpdateVmNetworkConfigParam UpdateVmNetworkConfig request param
type UpdateVmNetworkConfigParam struct {
	BaseParam
	Params UpdateVmNetworkConfigDetailParam `json:"params"`
}
