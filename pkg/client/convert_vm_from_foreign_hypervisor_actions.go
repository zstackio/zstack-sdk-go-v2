// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ConvertVmFromForeignHypervisor operates on ConvertVmFromForeignHypervisor
func (cli *ZSClient) ConvertVmFromForeignHypervisor(params param.ConvertVmFromForeignHypervisorParam) (*view.ConvertVmFromForeignHypervisorEventView, error) {
	resp := view.ConvertVmFromForeignHypervisorEventView{}
	if err := cli.Post("v1/v2vs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
