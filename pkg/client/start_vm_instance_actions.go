// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// StartVmInstance starts VmInstance
func (cli *ZSClient) StartVmInstance(uuid string, params param.StartVmInstanceParam) (*view.StartVmInstanceEventView, error) {
	resp := view.StartVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
