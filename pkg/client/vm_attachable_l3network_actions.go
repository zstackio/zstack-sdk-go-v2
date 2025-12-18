// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmAttachableL3Network 获取VmAttachableL3Network详情
func (cli *ZSClient) GetVmAttachableL3Network(uuid string) (*view.GetVmAttachableL3NetworkView, error) {
	var resp view.GetVmAttachableL3NetworkView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/l3-networks-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

