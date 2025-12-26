// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmCdRom creates VmCdRom
func (cli *ZSClient) CreateVmCdRom(params param.CreateVmCdRomParam) (*view.CreateVmCdRomEventView, error) {
	resp := view.CreateVmCdRomEventView{}
	if err := cli.Post("v1/vm-instances/cdroms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
