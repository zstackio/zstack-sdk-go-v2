// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmCdRom creates VmCdRom
func (cli *ZSClient) CreateVmCdRom(params param.CreateVmCdRomParam) (*view.CreateVmCdRomEventView, error) {
	resp := view.CreateVmCdRomEventView{}
	if err := cli.Post("v1/vm-instances/cdroms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
