// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdatePriorityConfig updates PriorityConfig
func (cli *ZSClient) UpdatePriorityConfig(uuid string, params param.UpdatePriorityConfigParam) (*view.UpdatePriorityConfigEventView, error) {
	resp := view.UpdatePriorityConfigEventView{}
	if err := cli.Put("v1/vm-priority-config/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
