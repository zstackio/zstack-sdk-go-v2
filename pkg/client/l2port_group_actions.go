// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateL2PortGroup 创建L2PortGroup
func (cli *ZSClient) CreateL2PortGroup(params param.CreateL2PortGroupParam) (*view.CreateL2PortGroupEventView, error) {
	resp := view.CreateL2PortGroupEventView{}
	if err := cli.Post("v1/l2-networks/port-group", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

