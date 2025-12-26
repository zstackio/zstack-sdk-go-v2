// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// FstrimVm operates on FstrimVm
func (cli *ZSClient) FstrimVm(params param.FstrimVmParam) (*view.FstrimVmEventView, error) {
	resp := view.FstrimVmEventView{}
	if err := cli.Post("v1/vm-instances/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
