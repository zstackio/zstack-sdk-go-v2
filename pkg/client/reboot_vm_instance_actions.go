// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RebootVmInstance operates on VmInstance
func (cli *ZSClient) RebootVmInstance(uuid string, params param.RebootVmInstanceParam) (*view.RebootVmInstanceEventView, error) {
	resp := view.RebootVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
