// Copyright (c) ZStack.io, Inc.

package param

// StartVmInstanceDetailParam StartVmInstance detail param
type StartVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
}

// StartVmInstanceParam StartVmInstance request param
type StartVmInstanceParam struct {
	BaseParam
	Params StartVmInstanceDetailParam `json:"params"`
}
