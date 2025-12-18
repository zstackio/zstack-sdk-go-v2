// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateBareMetal2Instance updates BareMetal2Instance
func (cli *ZSClient) UpdateBareMetal2Instance(uuid string, params param.UpdateBareMetal2InstanceParam) (*view.UpdateBareMetal2InstanceEventView, error) {
	resp := view.UpdateBareMetal2InstanceEventView{}
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/action", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
