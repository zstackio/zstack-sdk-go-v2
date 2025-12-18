// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdatePriorityConfig 更新PriorityConfig
func (cli *ZSClient) UpdatePriorityConfig(uuid string, params param.UpdatePriorityConfigParam) (*view.UpdatePriorityConfigEventView, error) {
	resp := view.UpdatePriorityConfigEventView{}
	if err := cli.Put("v1/vm-priority-config/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

