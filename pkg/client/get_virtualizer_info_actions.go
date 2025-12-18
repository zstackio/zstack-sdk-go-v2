// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVirtualizerInfo gets VirtualizerInfo by uuid
func (cli *ZSClient) GetVirtualizerInfo(uuid string) (*view.GetVirtualizerInfoView, error) {
	var resp view.GetVirtualizerInfoView
	if err := cli.Get("v1/vm-instances/virtualizer-info", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
