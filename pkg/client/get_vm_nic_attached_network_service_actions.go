// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmNicAttachedNetworkService gets VmNicAttachedNetworkService by uuid
func (cli *ZSClient) GetVmNicAttachedNetworkService(uuid string) (*view.GetVmNicAttachedNetworkServiceView, error) {
	var resp view.GetVmNicAttachedNetworkServiceView
	if err := cli.Get("v1/vm-instances/nics/{vmNicUuid}/attached-networkservices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
