// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVmCdRom updates VmCdRom
func (cli *ZSClient) UpdateVmCdRom(uuid string, params param.UpdateVmCdRomParam) (*view.UpdateVmCdRomEventView, error) {
	resp := view.UpdateVmCdRomEventView{}
	if err := cli.Put("v1/vm-instances/cdroms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
