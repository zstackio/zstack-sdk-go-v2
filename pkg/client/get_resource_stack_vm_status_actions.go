// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetResourceStackVmStatus gets ResourceStackVmStatus by uuid
func (cli *ZSClient) GetResourceStackVmStatus(uuid string) (*view.GetResourceStackVmStatusView, error) {
	var resp view.GetResourceStackVmStatusView
	if err := cli.Get("v1/cloudformation/stack/monitor/vmstatus", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
