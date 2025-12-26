// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmAttachableL3Network gets VmAttachableL3Network by uuid
func (cli *ZSClient) GetVmAttachableL3Network(uuid string) (*view.GetVmAttachableL3NetworkView, error) {
	var resp view.GetVmAttachableL3NetworkView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/l3-networks-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
