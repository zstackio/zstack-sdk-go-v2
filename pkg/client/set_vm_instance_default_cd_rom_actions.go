// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmInstanceDefaultCdRom operates on SetVmInstanceDefaultCdRom
func (cli *ZSClient) SetVmInstanceDefaultCdRom(uuid string, params param.SetVmInstanceDefaultCdRomParam) (*view.SetVmInstanceDefaultCdRomEventView, error) {
	resp := view.SetVmInstanceDefaultCdRomEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/cdroms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
