// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBareMetal2Instance creates BareMetal2Instance
func (cli *ZSClient) CreateBareMetal2Instance(params param.CreateBareMetal2InstanceParam) (*view.CreateBareMetal2InstanceEventView, error) {
	resp := view.CreateBareMetal2InstanceEventView{}
	if err := cli.Post("v1/baremetal2/bm-instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
