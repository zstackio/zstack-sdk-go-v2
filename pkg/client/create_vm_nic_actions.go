// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmNic creates VmNic
func (cli *ZSClient) CreateVmNic(params param.CreateVmNicParam) (*view.CreateVmNicEventView, error) {
	resp := view.CreateVmNicEventView{}
	if err := cli.Post("v1/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
