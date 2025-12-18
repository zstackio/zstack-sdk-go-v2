// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmNic 创建VmNic
func (cli *ZSClient) CreateVmNic(params param.CreateVmNicParam) (*view.CreateVmNicEventView, error) {
	resp := view.CreateVmNicEventView{}
	if err := cli.Post("v1/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

