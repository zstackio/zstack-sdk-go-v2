// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmInstanceDefaultCdRom 操作SetVmInstanceDefaultCdRom
func (cli *ZSClient) SetVmInstanceDefaultCdRom(uuid string, params param.SetVmInstanceDefaultCdRomParam) (*view.SetVmInstanceDefaultCdRomEventView, error) {
	resp := view.SetVmInstanceDefaultCdRomEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/cdroms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

