// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVmNicMac updates VmNicMac
func (cli *ZSClient) UpdateVmNicMac(uuid string, params param.UpdateVmNicMacParam) (*view.UpdateVmNicMacEventView, error) {
	resp := view.UpdateVmNicMacEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
