// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RecoverVmInstance operates on VmInstance
func (cli *ZSClient) RecoverVmInstance(uuid string, params param.RecoverVmInstanceParam) (*view.RecoverVmInstanceEventView, error) {
	resp := view.RecoverVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
