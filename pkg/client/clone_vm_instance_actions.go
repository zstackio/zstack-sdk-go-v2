// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CloneVmInstance operates on VmInstance
func (cli *ZSClient) CloneVmInstance(uuid string, params param.CloneVmInstanceParam) (*view.CloneVmInstanceEventView, error) {
	resp := view.CloneVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
